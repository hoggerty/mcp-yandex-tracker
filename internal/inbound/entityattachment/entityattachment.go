// Package entityattachment exposes Yandex Tracker entity attachment metadata operations as MCP tools.
package entityattachment

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentityattachment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityattachment"
)

func Tools(svc *logicentityattachment.Service) []server.ServerTool {
	tools := []tool.Tool{
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
