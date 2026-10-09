// Package component exposes Yandex Tracker queue-component operations as MCP tools.
package component

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
)

func Tools(svc *logiccomponent.Service) []server.ServerTool {
	tools := []tool.Tool{
		&listTool{svc: svc},
		&queueListTool{svc: svc},
		&getTool{svc: svc},
		&createTool{svc: svc},
		&editTool{svc: svc},
		&deleteTool{svc: svc},
		&accessUserGetTool{svc: svc},
		&accessGroupGetTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
