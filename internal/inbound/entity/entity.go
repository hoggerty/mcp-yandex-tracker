// Package entity exposes Yandex Tracker entity (project/portfolio/goal) operations as MCP tools.
package entity

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entity"
)

func Tools(svc *logicentity.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: logicentity.CreateSafety},
		{Tool: &getTool{svc: svc}, Safety: logicentity.GetSafety},
		{Tool: &updateTool{svc: svc}, Safety: logicentity.UpdateSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicentity.DeleteSafety},
		{Tool: &searchTool{svc: svc}, Safety: logicentity.SearchSafety},
		{Tool: &bulkUpdateTool{svc: svc}, Safety: logicentity.BulkUpdateSafety},
		{Tool: &eventsListTool{svc: svc}, Safety: logicentity.EventsListSafety},
	}
}
