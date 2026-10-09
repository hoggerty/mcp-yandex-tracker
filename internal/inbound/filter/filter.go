// Package filter exposes Yandex Tracker saved-filter operations as MCP tools.
package filter

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicfilter "github.com/hoggerty/yandex-tracker-mcp/internal/logic/filter"
)

func Tools(svc *logicfilter.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&getTool{svc: svc},
		&updateTool{svc: svc},
		&deleteTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
