// Package macro exposes Yandex Tracker queue-macro operations as MCP tools.
package macro

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicmacro "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/macro"
)

func Tools(svc *logicmacro.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
