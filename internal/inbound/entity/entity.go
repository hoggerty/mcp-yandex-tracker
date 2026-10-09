// Package entity exposes Yandex Tracker entity (project/portfolio/goal) operations as MCP tools.
package entity

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entity"
)

func Tools(svc *logicentity.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&getTool{svc: svc},
		&updateTool{svc: svc},
		&deleteTool{svc: svc},
		&searchTool{svc: svc},
		&bulkUpdateTool{svc: svc},
		&eventsListTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
