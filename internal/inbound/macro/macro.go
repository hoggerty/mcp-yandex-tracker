// Package macro exposes Yandex Tracker queue-macro operations as MCP tools.
package macro

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicmacro "github.com/hoggerty/yandex-tracker-mcp/internal/logic/macro"
)

func Tools(svc *logicmacro.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: logicmacro.ListSafety},
		{Tool: &getTool{svc: svc}, Safety: logicmacro.GetSafety},
		{Tool: &createTool{svc: svc}, Safety: logicmacro.CreateSafety},
		{Tool: &editTool{svc: svc}, Safety: logicmacro.EditSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicmacro.DeleteSafety},
	}
}
