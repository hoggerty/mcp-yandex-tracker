// Package macro exposes Yandex Tracker queue-macro operations as MCP tools.
package macro

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicmacro "github.com/hoggerty/yandex-tracker-mcp/internal/logic/macro"
)

func Tools(svc *logicmacro.Service) []server.ServerTool {
	tools := []tool.Tool{
		&listTool{svc: svc},
		&getTool{svc: svc},
		&createTool{svc: svc},
		&editTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
