// Package sprint exposes Yandex Tracker sprint operations as MCP tools.
package sprint

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicsprint "github.com/hoggerty/yandex-tracker-mcp/internal/logic/sprint"
)

func Tools(svc *logicsprint.Service) []server.ServerTool {
	tools := []tool.Tool{
		&getTool{svc: svc},
		&createTool{svc: svc},
		&editTool{svc: svc},
		&startTool{svc: svc},
		&archiveTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
