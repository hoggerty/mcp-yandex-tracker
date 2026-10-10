// Package entitylink exposes Yandex Tracker entity link operations as MCP tools.
package entitylink

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentitylink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitylink"
)

func Tools(svc *logicentitylink.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &addTool{svc: svc}, Safety: svc.AddSafety()},
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
