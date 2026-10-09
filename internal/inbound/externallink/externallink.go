// Package externallink exposes Yandex Tracker external-link operations as MCP tools.
package externallink

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicexternallink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/externallink"
)

func Tools(svc *logicexternallink.Service) []server.ServerTool {
	tools := []tool.Tool{
		&applicationsListTool{svc: svc},
		&listTool{svc: svc},
		&addTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
