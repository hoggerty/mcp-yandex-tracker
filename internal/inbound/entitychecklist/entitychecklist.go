// Package entitychecklist exposes Yandex Tracker entity checklist operations as MCP tools.
package entitychecklist

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentitychecklist "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitychecklist"
)

func Tools(svc *logicentitychecklist.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &addTool{svc: svc}, Safety: svc.AddSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &itemEditTool{svc: svc}, Safety: svc.ItemEditSafety()},
		{Tool: &itemMoveTool{svc: svc}, Safety: svc.ItemMoveSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &itemDeleteTool{svc: svc}, Safety: svc.ItemDeleteSafety()},
	}
}
