// Package report exposes Yandex Tracker issue-report operations as MCP tools.
package report

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicreport "github.com/hoggerty/yandex-tracker-mcp/internal/logic/report"
)

func Tools(svc *logicreport.Service) []server.ServerTool {
	tools := []tool.Tool{
		&createTool{svc: svc},
		&searchTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
