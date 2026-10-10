// Package bulkchange exposes Yandex Tracker bulk-change operations as MCP tools.
package bulkchange

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicbulkchange "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/bulkchange"
)

func Tools(svc *logicbulkchange.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &MoveTool{Svc: svc}, Safety: svc.MoveSafety()},
		{Tool: &UpdateTool{Svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &TransitionTool{Svc: svc}, Safety: svc.TransitionSafety()},
		{Tool: &StatusTool{Svc: svc}, Safety: svc.StatusSafety()},
		{Tool: &ErrorsTool{Svc: svc}, Safety: svc.ErrorsSafety()},
	}
}
