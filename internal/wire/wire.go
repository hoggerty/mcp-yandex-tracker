package wire

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound"
	"github.com/hoggerty/mcp-yandex-tracker/internal/logic"
	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
	"github.com/hoggerty/mcp-yandex-tracker/internal/outbound/trackerapi"
	"github.com/hoggerty/mcp-yandex-tracker/internal/wire/config"
)

// NewServer builds a Yandex Tracker MCP server with all tools registered.
// The returned server is transport-agnostic: callers decide how to serve it
// (e.g. server.ServeStdio, or a Streamable HTTP server).
func NewServer(version string) (*server.MCPServer, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	if err := cfg.TrackerClient.Validate(); err != nil {
		return nil, err
	}
	if err := cfg.Tools.Validate(); err != nil {
		return nil, err
	}

	client := trackerapi.New(trackerapi.Config{
		Token:      cfg.TrackerClient.Token,
		OrgID:      cfg.TrackerClient.OrgID,
		CloudOrgID: cfg.TrackerClient.CloudOrgID,
		BaseURL:    cfg.TrackerClient.BaseURL,
	})
	guard := safety.NewGuard(cfg.Tools.Mode)
	svcs := logic.New(client, guard)

	s := server.NewMCPServer("yandex-tracker-mcp", version)
	policy := inbound.Policy{
		Mode:         cfg.Tools.Mode,
		EnabledTools: cfg.Tools.Tools,
	}
	inbound.RegisterTools(s, svcs, policy)
	return s, nil
}
