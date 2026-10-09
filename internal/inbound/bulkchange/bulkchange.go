// Package bulkchange exposes Yandex Tracker bulk-change operations as MCP tools.
package bulkchange

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicbulkchange "github.com/hoggerty/yandex-tracker-mcp/internal/logic/bulkchange"
)

func Tools(svc *logicbulkchange.Service) []server.ServerTool {
	tools := []tool.Tool{
		&moveTool{svc: svc},
		&updateTool{svc: svc},
		&transitionTool{svc: svc},
		&statusTool{svc: svc},
		&errorsTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
