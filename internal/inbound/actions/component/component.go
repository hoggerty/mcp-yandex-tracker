// Package component exposes Yandex Tracker queue-component operations as MCP tools.
package component

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/component"
)

func Tools(svc *logiccomponent.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &QueueListTool{Svc: svc}, Safety: svc.QueueListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &AccessUserGetTool{Svc: svc}, Safety: svc.AccessUserGetSafety()},
		{Tool: &AccessGroupGetTool{Svc: svc}, Safety: svc.AccessGroupGetSafety()},
	}
}
