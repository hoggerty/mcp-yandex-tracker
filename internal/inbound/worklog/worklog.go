// Package worklog exposes Yandex Tracker worklog operations as MCP tools.
package worklog

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicworklog "github.com/hoggerty/yandex-tracker-mcp/internal/logic/worklog"
)

func Tools(svc *logicworklog.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &addTool{svc: svc}, Safety: logicworklog.AddSafety},
		{Tool: &editTool{svc: svc}, Safety: logicworklog.EditSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicworklog.DeleteSafety},
		{Tool: &listTool{svc: svc}, Safety: logicworklog.ListSafety},
		{Tool: &searchTool{svc: svc}, Safety: logicworklog.SearchSafety},
	}
}
