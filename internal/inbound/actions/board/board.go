// Package board exposes Yandex Tracker agile board operations as MCP tools.
package board

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicboard "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/board"
)

func Tools(svc *logicboard.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &SprintsListTool{Svc: svc}, Safety: svc.ListSprintsSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
