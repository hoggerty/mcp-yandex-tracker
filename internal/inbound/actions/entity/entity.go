// Package entity exposes Yandex Tracker entity (project/portfolio/goal) operations as MCP tools.
package entity

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicentity "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/entity"
)

func Tools(svc *logicentity.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &UpdateTool{Svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
		{Tool: &SearchTool{Svc: svc}, Safety: svc.SearchSafety()},
		{Tool: &BulkUpdateTool{Svc: svc}, Safety: svc.BulkUpdateSafety()},
		{Tool: &EventsListTool{Svc: svc}, Safety: svc.EventsListSafety()},
	}
}
