// Package bulkchange exposes Yandex Tracker bulk-change operations as MCP tools.
package bulkchange

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicbulkchange "github.com/hoggerty/yandex-tracker-mcp/internal/logic/bulkchange"
)

func Tools(svc *logicbulkchange.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &moveTool{svc: svc}, Safety: logicbulkchange.MoveSafety},
		{Tool: &updateTool{svc: svc}, Safety: logicbulkchange.UpdateSafety},
		{Tool: &transitionTool{svc: svc}, Safety: logicbulkchange.TransitionSafety},
		{Tool: &statusTool{svc: svc}, Safety: logicbulkchange.StatusSafety},
		{Tool: &errorsTool{svc: svc}, Safety: logicbulkchange.ErrorsSafety},
	}
}
