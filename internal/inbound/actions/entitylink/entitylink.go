// Package entitylink exposes Yandex Tracker entity link operations as MCP tools.
package entitylink

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentitylink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entitylink"
)

func Tools(svc *logicentitylink.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &AddTool{Svc: svc}, Safety: svc.AddSafety()},
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
