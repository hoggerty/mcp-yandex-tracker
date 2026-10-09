// Package field exposes Yandex Tracker global-field operations as MCP tools.
package field

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicfield "github.com/hoggerty/yandex-tracker-mcp/internal/logic/field"
)

func Tools(svc *logicfield.Service) []server.ServerTool {
	tools := []tool.Tool{
		&listTool{svc: svc},
		&createTool{svc: svc},
		&getTool{svc: svc},
		&editNameTool{svc: svc},
		&editValuesTool{svc: svc},
		&categoryCreateTool{svc: svc},
		&categoryEditTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
