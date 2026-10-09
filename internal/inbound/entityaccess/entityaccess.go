// Package entityaccess exposes Yandex Tracker entity access/permissions operations as MCP tools.
package entityaccess

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentityaccess "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityaccess"
)

func Tools(svc *logicentityaccess.Service) []server.ServerTool {
	tools := []tool.Tool{
		&accessGetTool{svc: svc},
		&permissionsGetTool{svc: svc},
		&accessEditTool{svc: svc},
		&permissionsEditTool{svc: svc},
	}

	out := make([]server.ServerTool, len(tools))
	for i, t := range tools {
		out[i] = tool.NewServerTool(t)
	}
	return out
}
