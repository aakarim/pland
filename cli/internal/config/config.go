package config

type ServerConfig struct {
	Host        string
	HttpScheme  string
	GraphQLPort int
	GraphQLPath string
}

type Config struct {
	ManagedPath string
	Server      ServerConfig
}

type ConfigFunc func(*Config)

func NewConfig(opts ...ConfigFunc) *Config {
	c := &Config{
		Server: ServerConfig{
			Host:        "localhost",
			HttpScheme:  "http",
			GraphQLPort: 8080,
			GraphQLPath: "/query",
		},
		ManagedPath: "",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func SetServer(s ServerConfig) ConfigFunc {
	return func(c *Config) {
		c.Server = s
	}
}
