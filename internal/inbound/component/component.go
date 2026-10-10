// Package component exposes Yandex Tracker queue-component operations as MCP tools.
package component

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
)

func Tools(svc *logiccomponent.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &queueListTool{svc: svc}, Safety: svc.QueueListSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &accessUserGetTool{svc: svc}, Safety: svc.AccessUserGetSafety()},
		{Tool: &accessGroupGetTool{svc: svc}, Safety: svc.AccessGroupGetSafety()},
	}
}
