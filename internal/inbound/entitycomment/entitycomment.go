// Package entitycomment exposes Yandex Tracker entity comment operations as MCP tools.
package entitycomment

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentitycomment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitycomment"
)

func Tools(svc *logicentitycomment.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &addTool{svc: svc}, Safety: svc.AddSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
