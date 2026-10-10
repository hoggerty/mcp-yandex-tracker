// Package column exposes Yandex Tracker agile board column operations as MCP tools.
package column

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logiccolumn "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/column"
)

func Tools(svc *logiccolumn.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
