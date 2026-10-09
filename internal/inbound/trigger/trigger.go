// Package trigger exposes Yandex Tracker queue-trigger operations as MCP tools.
package trigger

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logictrigger "github.com/hoggerty/yandex-tracker-mcp/internal/logic/trigger"
)

func Tools(svc *logictrigger.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&listTool{svc: svc},
		&getTool{svc: svc},
		&editTool{svc: svc},
		&logsListTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
