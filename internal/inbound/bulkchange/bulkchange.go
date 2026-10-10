// Package bulkchange exposes Yandex Tracker bulk-change operations as MCP tools.
package bulkchange

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicbulkchange "github.com/hoggerty/yandex-tracker-mcp/internal/logic/bulkchange"
)

func Tools(svc *logicbulkchange.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &moveTool{svc: svc}, Safety: svc.MoveSafety()},
		{Tool: &updateTool{svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &transitionTool{svc: svc}, Safety: svc.TransitionSafety()},
		{Tool: &statusTool{svc: svc}, Safety: svc.StatusSafety()},
		{Tool: &errorsTool{svc: svc}, Safety: svc.ErrorsSafety()},
	}
}
