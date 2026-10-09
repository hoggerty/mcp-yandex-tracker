// Package issuetype exposes Yandex Tracker issue type operations as MCP tools.
package issuetype

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicissuetype "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issuetype"
)

func Tools(svc *logicissuetype.Service) []server.ServerTool {
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
