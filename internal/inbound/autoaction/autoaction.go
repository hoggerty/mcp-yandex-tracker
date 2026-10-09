// Package autoaction exposes Yandex Tracker queue auto-action operations as MCP tools.
package autoaction

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicautoaction "github.com/hoggerty/yandex-tracker-mcp/internal/logic/autoaction"
)

func Tools(svc *logicautoaction.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&getTool{svc: svc},
		&logsListTool{svc: svc},
		&logGetTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
