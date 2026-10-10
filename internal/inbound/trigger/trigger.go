// Package trigger exposes Yandex Tracker queue-trigger operations as MCP tools.
package trigger

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logictrigger "github.com/hoggerty/yandex-tracker-mcp/internal/logic/trigger"
)

func Tools(svc *logictrigger.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &logsListTool{svc: svc}, Safety: svc.LogsListSafety()},
	}
}
