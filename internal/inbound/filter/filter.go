// Package filter exposes Yandex Tracker saved-filter operations as MCP tools.
package filter

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicfilter "github.com/hoggerty/yandex-tracker-mcp/internal/logic/filter"
)

func Tools(svc *logicfilter.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: logicfilter.CreateSafety},
		{Tool: &getTool{svc: svc}, Safety: logicfilter.GetSafety},
		{Tool: &updateTool{svc: svc}, Safety: logicfilter.UpdateSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicfilter.DeleteSafety},
	}
}
