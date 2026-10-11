// Package workflow exposes Yandex Tracker workflow operations as MCP tools.
package workflow

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicworkflow "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/workflow"
)

func Tools(svc *logicworkflow.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &QueueListTool{Svc: svc}, Safety: svc.QueueListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &ActionEditTool{Svc: svc}, Safety: svc.ActionEditSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
