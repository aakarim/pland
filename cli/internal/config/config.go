package config

import (
	"fmt"
	"io"
)

type ServerConfig struct {
	Host        string
	HttpScheme  string
	GraphQLPort int
	GraphQLPath string
}

type Config struct {
	ManagedPath   string
	PlanPath      string
	Server        ServerConfig
	EditorCommand string
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
		if opt == nil {
			continue
		}
		opt(c)
	}
	return c
}

func SetServer(s ServerConfig) ConfigFunc {
	return func(c *Config) {
		c.Server = s
	}
}

func SetPlanPath(p string) ConfigFunc {
	return func(c *Config) {
		c.PlanPath = p
	}
}

func WithFileValues(r io.Reader) (ConfigFunc, error) {
	fc, err := ParseFile(r)
	if err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	return func(c *Config) {
		c.EditorCommand = fc.EditorCommand
	}, nil
}
