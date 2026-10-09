// Package column exposes Yandex Tracker agile board column operations as MCP tools.
package column

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logiccolumn "github.com/hoggerty/yandex-tracker-mcp/internal/logic/column"
)

func Tools(svc *logiccolumn.Service) []server.ServerTool {
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
