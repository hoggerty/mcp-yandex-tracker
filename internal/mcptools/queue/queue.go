// Package queue exposes Yandex Tracker queue operations as MCP tools.
package queue

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/tool"
	"github.com/hoggerty/yandex-tracker-mcp/internal/trackerapi"
)

// Tools returns the MCP tools for Yandex Tracker queue operations.
func Tools(c *trackerapi.Client) []server.ServerTool {
	tools := []tool.Tool{
		&listTool{client: c},
		&getTool{client: c},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
