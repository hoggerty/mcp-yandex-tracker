// Package entityattachment exposes Yandex Tracker entity attachment metadata operations as MCP tools.
package entityattachment

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentityattachment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityattachment"
)

func Tools(svc *logicentityattachment.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: logicentityattachment.ListSafety},
		{Tool: &getTool{svc: svc}, Safety: logicentityattachment.GetSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicentityattachment.DeleteSafety},
	}
}
