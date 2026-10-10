// Package worklog exposes Yandex Tracker worklog operations as MCP tools.
package worklog

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicworklog "github.com/hoggerty/yandex-tracker-mcp/internal/logic/worklog"
)

func Tools(svc *logicworklog.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &addTool{svc: svc}, Safety: svc.AddSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &searchTool{svc: svc}, Safety: svc.SearchSafety()},
	}
}
