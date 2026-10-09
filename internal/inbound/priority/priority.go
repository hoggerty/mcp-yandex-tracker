// Package priority exposes Yandex Tracker issue priority operations as MCP tools.
package priority

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicpriority "github.com/hoggerty/yandex-tracker-mcp/internal/logic/priority"
)

func Tools(svc *logicpriority.Service) []server.ServerTool {
	tools := []tool.Tool{
		&listTool{svc: svc},
		&createTool{svc: svc},
		&editTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
