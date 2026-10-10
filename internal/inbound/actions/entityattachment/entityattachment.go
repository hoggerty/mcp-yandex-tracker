// Package entityattachment exposes Yandex Tracker entity attachment metadata operations as MCP tools.
package entityattachment

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentityattachment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entityattachment"
)

func Tools(svc *logicentityattachment.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
