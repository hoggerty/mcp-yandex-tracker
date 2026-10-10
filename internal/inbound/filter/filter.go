// Package filter exposes Yandex Tracker saved-filter operations as MCP tools.
package filter

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicfilter "github.com/hoggerty/yandex-tracker-mcp/internal/logic/filter"
)

func Tools(svc *logicfilter.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &updateTool{svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
