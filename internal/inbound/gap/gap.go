// Package gap exposes Yandex Tracker employee absence (gap) operations as MCP tools.
package gap

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicgap "github.com/hoggerty/yandex-tracker-mcp/internal/logic/gap"
)

func Tools(svc *logicgap.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &searchTool{svc: svc}, Safety: svc.SearchSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
