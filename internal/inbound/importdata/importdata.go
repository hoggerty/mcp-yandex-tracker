// Package importdata exposes Yandex Tracker historical-data import operations as MCP tools.
package importdata

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicimportdata "github.com/hoggerty/yandex-tracker-mcp/internal/logic/importdata"
)

func Tools(svc *logicimportdata.Service) []server.ServerTool {
	tools := []tool.Tool{
		&ticketTool{svc: svc},
		&commentTool{svc: svc},
		&linkTool{svc: svc},
		&worklogTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
