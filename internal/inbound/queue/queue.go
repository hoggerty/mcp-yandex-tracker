// Package queue exposes Yandex Tracker queue operations as MCP tools.
package queue

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/queue"
)

func Tools(svc *logicqueue.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: logicqueue.ListSafety},
		{Tool: &getTool{svc: svc}, Safety: logicqueue.GetSafety},
		{Tool: &createTool{svc: svc}, Safety: logicqueue.CreateSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicqueue.DeleteSafety},
		{Tool: &restoreTool{svc: svc}, Safety: logicqueue.RestoreSafety},
		{Tool: &fieldsListTool{svc: svc}, Safety: logicqueue.FieldsListSafety},
		{Tool: &tagsListTool{svc: svc}, Safety: logicqueue.TagsListSafety},
		{Tool: &tagDeleteTool{svc: svc}, Safety: logicqueue.TagDeleteSafety},
		{Tool: &accessManageTool{svc: svc}, Safety: logicqueue.AccessManageSafety},
		{Tool: &accessUserGetTool{svc: svc}, Safety: logicqueue.AccessUserGetSafety},
		{Tool: &accessGroupGetTool{svc: svc}, Safety: logicqueue.AccessGroupGetSafety},
		{Tool: &localFieldCreateTool{svc: svc}, Safety: logicqueue.LocalFieldCreateSafety},
		{Tool: &localFieldsListTool{svc: svc}, Safety: logicqueue.LocalFieldsListSafety},
		{Tool: &localFieldGetTool{svc: svc}, Safety: logicqueue.LocalFieldGetSafety},
		{Tool: &localFieldEditTool{svc: svc}, Safety: logicqueue.LocalFieldEditSafety},
	}
}
