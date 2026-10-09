// Package worklog exposes Yandex Tracker worklog operations as MCP tools.
package worklog

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicworklog "github.com/hoggerty/yandex-tracker-mcp/internal/logic/worklog"
)

func Tools(svc *logicworklog.Service) []server.ServerTool {
	tools := []tool.Tool{
		&addTool{svc: svc},
		&editTool{svc: svc},
		&deleteTool{svc: svc},
		&listTool{svc: svc},
		&searchTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
