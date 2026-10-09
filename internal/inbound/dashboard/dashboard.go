// Package dashboard exposes Yandex Tracker dashboard operations as MCP tools.
package dashboard

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicdashboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/dashboard"
)

func Tools(svc *logicdashboard.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&cycleTimeWidgetCreateTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
