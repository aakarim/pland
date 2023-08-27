package main

import (
	"text/template"

	"entgo.io/contrib/entgql"
	gossh "golang.org/x/crypto/ssh"

	"context"
	"crypto/ed25519"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gliderlabs/ssh"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	_ "github.com/xiaoqidun/entps"
	"goji.io"
	"goji.io/pat"

	"github.com/aakarim/pland/ent"
	"github.com/aakarim/pland/ent/migrate"
	"github.com/aakarim/pland/ent/user"
	"github.com/aakarim/pland/graph"
	"github.com/aakarim/pland/graph/generated"
	"github.com/charmbracelet/charm/server"
	"github.com/charmbracelet/charm/server/db/sqlite"
	"github.com/charmbracelet/keygen"
)

func main() {
	cfg := server.DefaultConfig()
	// ent init
	sqliteDsn := fmt.Sprintf("file:./%s"+sqlite.DbOptions, filepath.Join(cfg.DataDir, "db", sqlite.DbName))
	log.Println("opening sqlite DB", sqliteDsn)
	client, err := ent.Open("sqlite3", sqliteDsn)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(),
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropColumn(true), // TODO: remove in prod
		migrate.WithDropIndex(true),  // TODO: remove in prod?
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	sp := filepath.Join(cfg.DataDir, ".ssh")
	kp, err := keygen.NewWithWrite(filepath.Join(sp, "charm_server"), []byte(""), keygen.Ed25519)
	if err != nil {
		log.Fatal(err)
	}
	cfg = cfg.WithKeys(kp.PublicKey(), kp.PrivateKeyPEM())
	s, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	gqlServer := makeGQLServer(client, cfg)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := s.Start(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Fatalf("error starting server: %s", err)
		}
	}()

	go func() {
		if err := gqlServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error starting server: %s", err)
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	gqlServer.Shutdown(ctx)
	s.Shutdown(ctx)
}

func makeGQLServer(client *ent.Client, cfg *server.Config) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	gqlMux := goji.NewMux()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{client}}))
	srv.Use(entgql.Transactioner{TxOpener: client})
	pk, err := gossh.ParseRawPrivateKey(cfg.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	jwtKeyPair := server.NewJSONWebKeyPair(pk.(*ed25519.PrivateKey))
	authMiddleware, err := AuthMiddleware(
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	jwtMiddleware, err := JWTMiddleware(
		jwtKeyPair.JWK.Public(),
		"http://localhost:35354",
		[]string{"charm"}, // TODO: what do we use here?
	)
	if err != nil {
		log.Fatal(err)
	}

	gqlMux.Use(jwtMiddleware)
	gqlMux.Use(authMiddleware)

	gqlMux.Handle(pat.New("/dev/playground"), playground.Handler("GraphQL playground", "/query"))
	gqlMux.Handle(pat.New("/query"), srv)
	// all the user handles
	gqlMux.Handle(pat.New("/*"), &UserHandler{entClient: client})

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: gqlMux,
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Println()
	return httpServer
}

type UserHandler struct {
	entClient *ent.Client
}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[1:]

	user, err := u.entClient.User.Query().Where(user.NameEQ(username)).Only(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("user %s not found", username), http.StatusNotFound)
		return
	}

	plan, err := user.QueryPlans().First(r.Context())
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "no plans yet!")
		return
	}

	t, err := template.New("user").Parse(userTmpl)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(plan.Txt))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	html := markdown.Render(doc, renderer)
	t.Execute(w, string(html))
}
