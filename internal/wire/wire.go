package wire

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic"
	"github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"
	"github.com/hoggerty/yandex-tracker-mcp/internal/wire/config"
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

	client := trackerapi.New(trackerapi.Config{
		Token:      cfg.TrackerClient.Token,
		OrgID:      cfg.TrackerClient.OrgID,
		CloudOrgID: cfg.TrackerClient.CloudOrgID,
		BaseURL:    cfg.TrackerClient.BaseURL,
	})
	svcs := logic.New(client)

	s := server.NewMCPServer("yandex-tracker-mcp", version)
	inbound.RegisterTools(s, svcs)
	return s, nil
}
