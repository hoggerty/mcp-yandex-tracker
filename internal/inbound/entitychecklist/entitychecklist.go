// Package entitychecklist exposes Yandex Tracker entity checklist operations as MCP tools.
package entitychecklist

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentitychecklist "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitychecklist"
)

func Tools(svc *logicentitychecklist.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &addTool{svc: svc}, Safety: logicentitychecklist.AddSafety},
		{Tool: &editTool{svc: svc}, Safety: logicentitychecklist.EditSafety},
		{Tool: &itemEditTool{svc: svc}, Safety: logicentitychecklist.ItemEditSafety},
		{Tool: &itemMoveTool{svc: svc}, Safety: logicentitychecklist.ItemMoveSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicentitychecklist.DeleteSafety},
		{Tool: &itemDeleteTool{svc: svc}, Safety: logicentitychecklist.ItemDeleteSafety},
	}
}
