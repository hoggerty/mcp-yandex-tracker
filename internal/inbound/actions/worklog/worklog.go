// Package worklog exposes Yandex Tracker worklog operations as MCP tools.
package worklog

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicworklog "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/worklog"
)

func Tools(svc *logicworklog.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &AddTool{Svc: svc}, Safety: svc.AddSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &SearchTool{Svc: svc}, Safety: svc.SearchSafety()},
	}
}
