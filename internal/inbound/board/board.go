// Package board exposes Yandex Tracker agile board operations as MCP tools.
package board

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/board"
)

func Tools(svc *logicboard.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &sprintsListTool{svc: svc}, Safety: svc.ListSprintsSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
