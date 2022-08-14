package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/kardianos/service"
)

const serviceName = "PlanD"
const serviceDescription = "The Plan Daemon. Watches your .plan files and syncs them with the server."

var svcFlag = flag.String("service", "", "Control the system service.")
var host = flag.String("host", "", "Host for server.")
var planExecPath = flag.String("planExecPath", "", "Path to the `plan` executable")

type program struct {
	watcher *fsnotify.Watcher
	baseCtx context.Context
	homeDir string
}

func (p program) Start(s service.Service) error {
	flag.Parse()
	log.Println("starting service", p.homeDir, os.Args[2], os.Args[4], os.Args, os.Environ())

	path := filepath.Join(p.homeDir, ".plan")

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-p.watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				log.Println("modified file:", event.Name)
				if event.Op == fsnotify.Write {
					c := exec.CommandContext(p.baseCtx, os.Args[4], "-d")
					c.Env = []string{
						"CHARM_HOST=" + os.Args[2],
						"HOME=" + p.homeDir,
					}
					c.Stdout = os.Stdout
					c.Stderr = os.Stderr
					if err := c.Run(); err != nil {
						log.Println("plan error:", err)
					}
				}
			case err, ok := <-p.watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err := p.watcher.Add(path)
	if err != nil {
		log.Fatalf("Watcher.Add(%s): %s", p.homeDir, err.Error())
	}

	// Block main goroutine forever.
	log.Println("watching", path)
	return nil
}

func (p program) Stop(s service.Service) error {
	log.Println("stopping service")
	return p.watcher.Close()
}

func main() {
	flag.Parse()
	ctx, shutdown := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	// get user to run on - we need the environment variables
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	// planExecPath, err := exec.LookPath("plan")
	// if err == nil {
	// 	planExecPath, err = filepath.Abs(planExecPath)
	// }
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	options := make(service.KeyValue)
	options["UserService"] = true // fix error 78
	options["RunAtLoad"] = true
	options["KeepAlive"] = true

	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		UserName:    u.Username,
		Arguments: []string{
			"host",
			*host,
			"planExecPath",
			*planExecPath,
		},
		Description: serviceDescription,
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
		Option: options,
	}

	// Create new watcher to be shared between different funcs
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	prg := &program{
		watcher: watcher, baseCtx: ctx, homeDir: homeDir,
	}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service: " + err.Error())
	}
	if len(*svcFlag) != 0 {
		log.Printf("running command: %s", *svcFlag)
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
	defer shutdown()
	<-ctx.Done()
}
