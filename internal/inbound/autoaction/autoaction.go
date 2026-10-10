// Package autoaction exposes Yandex Tracker queue auto-action operations as MCP tools.
package autoaction

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicautoaction "github.com/hoggerty/yandex-tracker-mcp/internal/logic/autoaction"
)

func Tools(svc *logicautoaction.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &logsListTool{svc: svc}, Safety: svc.LogsListSafety()},
		{Tool: &logGetTool{svc: svc}, Safety: svc.LogGetSafety()},
	}
}
