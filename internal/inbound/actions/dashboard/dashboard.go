// Package dashboard exposes Yandex Tracker dashboard operations as MCP tools.
package dashboard

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicdashboard "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/dashboard"
)

func Tools(svc *logicdashboard.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &CycleTimeWidgetCreateTool{Svc: svc}, Safety: svc.CreateCycleTimeWidgetSafety()},
	}
}
