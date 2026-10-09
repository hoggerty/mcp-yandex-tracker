// Package component exposes Yandex Tracker queue-component operations as MCP tools.
package component

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
)

func Tools(svc *logiccomponent.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: logiccomponent.ListSafety},
		{Tool: &queueListTool{svc: svc}, Safety: logiccomponent.QueueListSafety},
		{Tool: &getTool{svc: svc}, Safety: logiccomponent.GetSafety},
		{Tool: &createTool{svc: svc}, Safety: logiccomponent.CreateSafety},
		{Tool: &editTool{svc: svc}, Safety: logiccomponent.EditSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logiccomponent.DeleteSafety},
		{Tool: &accessUserGetTool{svc: svc}, Safety: logiccomponent.AccessUserGetSafety},
		{Tool: &accessGroupGetTool{svc: svc}, Safety: logiccomponent.AccessGroupGetSafety},
	}
}
