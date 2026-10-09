// Package status exposes Yandex Tracker issue status operations as MCP tools.
package status

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicstatus "github.com/hoggerty/yandex-tracker-mcp/internal/logic/status"
)

func Tools(svc *logicstatus.Service) []server.ServerTool {
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
