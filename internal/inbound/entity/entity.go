// Package entity exposes Yandex Tracker entity (project/portfolio/goal) operations as MCP tools.
package entity

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entity"
)

func Tools(svc *logicentity.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &updateTool{svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &searchTool{svc: svc}, Safety: svc.SearchSafety()},
		{Tool: &bulkUpdateTool{svc: svc}, Safety: svc.BulkUpdateSafety()},
		{Tool: &eventsListTool{svc: svc}, Safety: svc.EventsListSafety()},
	}
}
