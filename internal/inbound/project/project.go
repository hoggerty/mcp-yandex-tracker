// Package project exposes Yandex Tracker project operations as MCP tools.
package project

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicproject "github.com/hoggerty/yandex-tracker-mcp/internal/logic/project"
)

func Tools(svc *logicproject.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&getTool{svc: svc},
		&listTool{svc: svc},
		&queuesListTool{svc: svc},
		&updateTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
