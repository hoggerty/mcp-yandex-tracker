// Package workflow exposes Yandex Tracker workflow operations as MCP tools.
package workflow

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicworkflow "github.com/hoggerty/yandex-tracker-mcp/internal/logic/workflow"
)

func Tools(svc *logicworkflow.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: logicworkflow.CreateSafety},
		{Tool: &listTool{svc: svc}, Safety: logicworkflow.ListSafety},
		{Tool: &queueListTool{svc: svc}, Safety: logicworkflow.QueueListSafety},
		{Tool: &getTool{svc: svc}, Safety: logicworkflow.GetSafety},
		{Tool: &editTool{svc: svc}, Safety: logicworkflow.EditSafety},
		{Tool: &actionEditTool{svc: svc}, Safety: logicworkflow.ActionEditSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicworkflow.DeleteSafety},
	}
}
