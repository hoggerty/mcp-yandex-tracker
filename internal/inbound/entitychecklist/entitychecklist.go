// Package entitychecklist exposes Yandex Tracker entity checklist operations as MCP tools.
package entitychecklist

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentitychecklist "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitychecklist"
)

func Tools(svc *logicentitychecklist.Service) []server.ServerTool {
	tools := []tool.Tool{
		&addTool{svc: svc},
		&editTool{svc: svc},
		&itemEditTool{svc: svc},
		&itemMoveTool{svc: svc},
		&deleteTool{svc: svc},
		&itemDeleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
