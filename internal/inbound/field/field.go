// Package field exposes Yandex Tracker global-field operations as MCP tools.
package field

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicfield "github.com/hoggerty/yandex-tracker-mcp/internal/logic/field"
)

func Tools(svc *logicfield.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &editNameTool{svc: svc}, Safety: svc.EditNameSafety()},
		{Tool: &editValuesTool{svc: svc}, Safety: svc.EditValuesSafety()},
		{Tool: &categoryCreateTool{svc: svc}, Safety: svc.CategoryCreateSafety()},
		{Tool: &categoryEditTool{svc: svc}, Safety: svc.CategoryEditSafety()},
	}
}
