// Package version exposes Yandex Tracker queue-version operations as MCP tools.
package version

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/version"
)

func Tools(svc *logicversion.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&queueListTool{svc: svc},
		&getTool{svc: svc},
		&editTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
