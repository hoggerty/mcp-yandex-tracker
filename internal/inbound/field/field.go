// Package field exposes Yandex Tracker global-field operations as MCP tools.
package field

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicfield "github.com/hoggerty/yandex-tracker-mcp/internal/logic/field"
)

func Tools(svc *logicfield.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: logicfield.ListSafety},
		{Tool: &createTool{svc: svc}, Safety: logicfield.CreateSafety},
		{Tool: &getTool{svc: svc}, Safety: logicfield.GetSafety},
		{Tool: &editNameTool{svc: svc}, Safety: logicfield.EditNameSafety},
		{Tool: &editValuesTool{svc: svc}, Safety: logicfield.EditValuesSafety},
		{Tool: &categoryCreateTool{svc: svc}, Safety: logicfield.CategoryCreateSafety},
		{Tool: &categoryEditTool{svc: svc}, Safety: logicfield.CategoryEditSafety},
	}
}
