// Package queue exposes Yandex Tracker queue operations as MCP tools.
package queue

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/queue"
)

func Tools(svc *logicqueue.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &restoreTool{svc: svc}, Safety: svc.RestoreSafety()},
		{Tool: &fieldsListTool{svc: svc}, Safety: svc.FieldsListSafety()},
		{Tool: &tagsListTool{svc: svc}, Safety: svc.TagsListSafety()},
		{Tool: &tagDeleteTool{svc: svc}, Safety: svc.TagDeleteSafety()},
		{Tool: &accessManageTool{svc: svc}, Safety: svc.AccessManageSafety()},
		{Tool: &accessUserGetTool{svc: svc}, Safety: svc.AccessUserGetSafety()},
		{Tool: &accessGroupGetTool{svc: svc}, Safety: svc.AccessGroupGetSafety()},
		{Tool: &localFieldCreateTool{svc: svc}, Safety: svc.LocalFieldCreateSafety()},
		{Tool: &localFieldsListTool{svc: svc}, Safety: svc.LocalFieldsListSafety()},
		{Tool: &localFieldGetTool{svc: svc}, Safety: svc.LocalFieldGetSafety()},
		{Tool: &localFieldEditTool{svc: svc}, Safety: svc.LocalFieldEditSafety()},
	}
}
