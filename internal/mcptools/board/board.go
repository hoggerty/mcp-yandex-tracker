// Package board exposes Yandex Tracker agile board operations as MCP tools.
package board

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/tool"
	"github.com/hoggerty/yandex-tracker-mcp/internal/trackerapi"
)

// Tools returns the MCP tools for Yandex Tracker agile board operations.
func Tools(c *trackerapi.Client) []server.ServerTool {
	tools := []tool.Tool{
		&listTool{client: c},
		&sprintsListTool{client: c},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
