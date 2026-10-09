// Package board exposes Yandex Tracker agile board operations as MCP tools.
package board

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/board"
)

func Tools(svc *logicboard.Service) []server.ServerTool {
	tools := []tool.Tool{
		&listTool{svc: svc},
		&sprintsListTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
