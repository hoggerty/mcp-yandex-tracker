// Package user exposes Yandex Tracker user operations as MCP tools.
package user

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicuser "github.com/hoggerty/yandex-tracker-mcp/internal/logic/user"
)

func Tools(svc *logicuser.Service) []server.ServerTool {
	tools := []tool.Tool{
		&currentTool{svc: svc},
		&listTool{svc: svc},
		&getTool{svc: svc},
		&listRelativeTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
