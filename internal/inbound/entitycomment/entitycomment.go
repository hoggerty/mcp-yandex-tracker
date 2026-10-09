// Package entitycomment exposes Yandex Tracker entity comment operations as MCP tools.
package entitycomment

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentitycomment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitycomment"
)

func Tools(svc *logicentitycomment.Service) []server.ServerTool {
	tools := []tool.Tool{
		&addTool{svc: svc},
		&editTool{svc: svc},
		&listTool{svc: svc},
		&getTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
