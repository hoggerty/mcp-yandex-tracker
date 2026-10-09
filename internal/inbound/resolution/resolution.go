// Package resolution exposes Yandex Tracker issue resolution operations as MCP tools.
package resolution

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicresolution "github.com/hoggerty/yandex-tracker-mcp/internal/logic/resolution"
)

func Tools(svc *logicresolution.Service) []server.ServerTool {
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
