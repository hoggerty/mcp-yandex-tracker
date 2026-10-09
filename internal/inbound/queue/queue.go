// Package queue exposes Yandex Tracker queue operations as MCP tools.
package queue

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/queue"
)

func Tools(svc *logicqueue.Service) []server.ServerTool {
	tools := []tool.Tool{
		&listTool{svc: svc},
		&getTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
