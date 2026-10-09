// Package issue exposes Yandex Tracker issue operations as MCP tools.
package issue

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

func Tools(svc *logicissue.Service) []server.ServerTool {
	tools := []tool.Tool{
		&getTool{svc: svc},
		&searchTool{svc: svc},
		&countTool{svc: svc},
		&createTool{svc: svc},
		&updateTool{svc: svc},
		&transitionsListTool{svc: svc},
		&transitionExecuteTool{svc: svc},
		&commentsListTool{svc: svc},
		&commentAddTool{svc: svc},
		&linksListTool{svc: svc},
		&linkAddTool{svc: svc},
		&changelogGetTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
