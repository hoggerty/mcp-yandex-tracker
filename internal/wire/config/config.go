// Package config sources the configuration needed to wire the Yandex
// Tracker MCP server. It is the only package that knows configuration comes
// from environment variables — outbound, logic and inbound never read the
// environment themselves.
package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	TrackerClient TrackerClientConfig
	Tools         ToolsConfig
}

// Load reads Config from environment variables, without validating it.
func Load() (Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}
	return cfg, nil
}
