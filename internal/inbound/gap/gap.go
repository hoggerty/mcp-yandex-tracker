// Package gap exposes Yandex Tracker employee absence (gap) operations as MCP tools.
package gap

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicgap "github.com/hoggerty/yandex-tracker-mcp/internal/logic/gap"
)

func Tools(svc *logicgap.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&searchTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
