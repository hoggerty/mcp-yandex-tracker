// Package column exposes Yandex Tracker agile board column operations as MCP tools.
package column

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logiccolumn "github.com/hoggerty/yandex-tracker-mcp/internal/logic/column"
)

func Tools(svc *logiccolumn.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
