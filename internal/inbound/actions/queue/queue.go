// Package queue exposes Yandex Tracker queue operations as MCP tools.
package queue

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicqueue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/queue"
)

func Tools(svc *logicqueue.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &RestoreTool{Svc: svc}, Safety: svc.RestoreSafety()},
		{Tool: &FieldsListTool{Svc: svc}, Safety: svc.FieldsListSafety()},
		{Tool: &TagsListTool{Svc: svc}, Safety: svc.TagsListSafety()},
		{Tool: &TagDeleteTool{Svc: svc}, Safety: svc.TagDeleteSafety()},
		{Tool: &AccessManageTool{Svc: svc}, Safety: svc.AccessManageSafety()},
		{Tool: &AccessUserGetTool{Svc: svc}, Safety: svc.AccessUserGetSafety()},
		{Tool: &AccessGroupGetTool{Svc: svc}, Safety: svc.AccessGroupGetSafety()},
		{Tool: &LocalFieldCreateTool{Svc: svc}, Safety: svc.LocalFieldCreateSafety()},
		{Tool: &LocalFieldsListTool{Svc: svc}, Safety: svc.LocalFieldsListSafety()},
		{Tool: &LocalFieldGetTool{Svc: svc}, Safety: svc.LocalFieldGetSafety()},
		{Tool: &LocalFieldEditTool{Svc: svc}, Safety: svc.LocalFieldEditSafety()},
	}
}
