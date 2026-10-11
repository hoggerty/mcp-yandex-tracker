// Package autoaction exposes Yandex Tracker queue auto-action operations as MCP tools.
package autoaction

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicautoaction "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/autoaction"
)

func Tools(svc *logicautoaction.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &LogsListTool{Svc: svc}, Safety: svc.LogsListSafety()},
		{Tool: &LogGetTool{Svc: svc}, Safety: svc.LogGetSafety()},
	}
}
