// Package dashboard exposes Yandex Tracker dashboard operations as MCP tools.
package dashboard

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicdashboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/dashboard"
)

func Tools(svc *logicdashboard.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &cycleTimeWidgetCreateTool{svc: svc}, Safety: svc.CreateCycleTimeWidgetSafety()},
	}
}
