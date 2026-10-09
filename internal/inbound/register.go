package inbound

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/board"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/issue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/queue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/user"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic"
)

func RegisterTools(s *server.MCPServer, svcs *logic.Services) {
	s.AddTools(allTools(svcs)...)
}

func allTools(svcs *logic.Services) []server.ServerTool {
	var all []server.ServerTool
	all = append(all, issue.Tools(svcs.Issue)...)
	all = append(all, queue.Tools(svcs.Queue)...)
	all = append(all, user.Tools(svcs.User)...)
	all = append(all, board.Tools(svcs.Board)...)
	return all
}
