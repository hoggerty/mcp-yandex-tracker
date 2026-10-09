// Package trigger exposes Yandex Tracker queue-trigger operations as MCP tools.
package trigger

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logictrigger "github.com/hoggerty/yandex-tracker-mcp/internal/logic/trigger"
)

func Tools(svc *logictrigger.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: logictrigger.CreateSafety},
		{Tool: &listTool{svc: svc}, Safety: logictrigger.ListSafety},
		{Tool: &getTool{svc: svc}, Safety: logictrigger.GetSafety},
		{Tool: &editTool{svc: svc}, Safety: logictrigger.EditSafety},
		{Tool: &logsListTool{svc: svc}, Safety: logictrigger.LogsListSafety},
	}
}
