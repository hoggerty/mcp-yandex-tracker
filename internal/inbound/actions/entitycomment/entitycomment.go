// Package entitycomment exposes Yandex Tracker entity comment operations as MCP tools.
package entitycomment

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicentitycomment "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/entitycomment"
)

func Tools(svc *logicentitycomment.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &AddTool{Svc: svc}, Safety: svc.AddSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
