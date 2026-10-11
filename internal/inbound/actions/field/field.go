// Package field exposes Yandex Tracker global-field operations as MCP tools.
package field

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicfield "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/field"
)

func Tools(svc *logicfield.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &EditNameTool{Svc: svc}, Safety: svc.EditNameSafety()},
		{Tool: &EditValuesTool{Svc: svc}, Safety: svc.EditValuesSafety()},
		{Tool: &CategoryCreateTool{Svc: svc}, Safety: svc.CategoryCreateSafety()},
		{Tool: &CategoryEditTool{Svc: svc}, Safety: svc.CategoryEditSafety()},
	}
}
