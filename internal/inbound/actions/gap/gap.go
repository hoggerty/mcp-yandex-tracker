// Package gap exposes Yandex Tracker employee absence (gap) operations as MCP tools.
package gap

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicgap "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/gap"
)

func Tools(svc *logicgap.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &SearchTool{Svc: svc}, Safety: svc.SearchSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
