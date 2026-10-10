// Package entitychecklist exposes Yandex Tracker entity checklist operations as MCP tools.
package entitychecklist

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentitychecklist "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entitychecklist"
)

func Tools(svc *logicentitychecklist.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &AddTool{Svc: svc}, Safety: svc.AddSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &ItemEditTool{Svc: svc}, Safety: svc.ItemEditSafety()},
		{Tool: &ItemMoveTool{Svc: svc}, Safety: svc.ItemMoveSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &ItemDeleteTool{Svc: svc}, Safety: svc.ItemDeleteSafety()},
	}
}
