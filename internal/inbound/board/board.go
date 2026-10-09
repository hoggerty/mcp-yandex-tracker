// Package board exposes Yandex Tracker agile board operations as MCP tools.
package board

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/board"
)

func Tools(svc *logicboard.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: logicboard.ListSafety},
		{Tool: &sprintsListTool{svc: svc}, Safety: logicboard.ListSprintsSafety},
		{Tool: &getTool{svc: svc}, Safety: logicboard.GetSafety},
		{Tool: &createTool{svc: svc}, Safety: logicboard.CreateSafety},
		{Tool: &editTool{svc: svc}, Safety: logicboard.EditSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicboard.DeleteSafety},
	}
}
