// Package workflow exposes Yandex Tracker workflow operations as MCP tools.
package workflow

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicworkflow "github.com/hoggerty/yandex-tracker-mcp/internal/logic/workflow"
)

func Tools(svc *logicworkflow.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &queueListTool{svc: svc}, Safety: svc.QueueListSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &actionEditTool{svc: svc}, Safety: svc.ActionEditSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
