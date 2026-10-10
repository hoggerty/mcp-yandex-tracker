// Package filter exposes Yandex Tracker saved-filter operations as MCP tools.
package filter

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicfilter "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/filter"
)

func Tools(svc *logicfilter.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &UpdateTool{Svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
