package mcptools

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/board"
	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/issue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/queue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/user"
	"github.com/hoggerty/yandex-tracker-mcp/internal/trackerapi"
)

// RegisterTools adds all Yandex Tracker tools to the given MCP server.
func RegisterTools(s *server.MCPServer, client *trackerapi.Client) {
	s.AddTools(allTools(client)...)
}

func allTools(c *trackerapi.Client) []server.ServerTool {
	var all []server.ServerTool
	all = append(all, issue.Tools(c)...)
	all = append(all, queue.Tools(c)...)
	all = append(all, user.Tools(c)...)
	all = append(all, board.Tools(c)...)
	return all
}
