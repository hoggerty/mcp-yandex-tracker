// Package workflow exposes Yandex Tracker workflow operations as MCP tools.
package workflow

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicworkflow "github.com/hoggerty/yandex-tracker-mcp/internal/logic/workflow"
)

func Tools(svc *logicworkflow.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&listTool{svc: svc},
		&queueListTool{svc: svc},
		&getTool{svc: svc},
		&editTool{svc: svc},
		&actionEditTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
