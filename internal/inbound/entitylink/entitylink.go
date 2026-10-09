// Package entitylink exposes Yandex Tracker entity link operations as MCP tools.
package entitylink

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentitylink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitylink"
)

func Tools(svc *logicentitylink.Service) []server.ServerTool {
	tools := []tool.Tool{
		&addTool{svc: svc},
		&listTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
