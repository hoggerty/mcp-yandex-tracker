package wire

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic"
	"github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"
)

// NewServer builds a Yandex Tracker MCP server with all tools registered.
// The returned server is transport-agnostic: callers decide how to serve it
// (e.g. server.ServeStdio, or a Streamable HTTP server).
func NewServer(cfg trackerapi.Config, version string) *server.MCPServer {
	client := trackerapi.New(cfg)
	svcs := logic.New(client)

	s := server.NewMCPServer("yandex-tracker-mcp", version)
	inbound.RegisterTools(s, svcs)
	return s
}
