// Package trigger exposes Yandex Tracker queue-trigger operations as MCP tools.
package trigger

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logictrigger "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/trigger"
)

func Tools(svc *logictrigger.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &LogsListTool{Svc: svc}, Safety: svc.LogsListSafety()},
	}
}
