// Package issue exposes Yandex Tracker issue operations as MCP tools.
package issue

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/tool"
	"github.com/hoggerty/yandex-tracker-mcp/internal/trackerapi"
)

// Tools returns the MCP tools for Yandex Tracker issue operations.
func Tools(c *trackerapi.Client) []server.ServerTool {
	tools := []tool.Tool{
		&getTool{client: c},
		&searchTool{client: c},
		&countTool{client: c},
		&createTool{client: c},
		&updateTool{client: c},
		&transitionsListTool{client: c},
		&transitionExecuteTool{client: c},
		&commentsListTool{client: c},
		&commentAddTool{client: c},
		&linksListTool{client: c},
		&linkAddTool{client: c},
		&changelogGetTool{client: c},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
