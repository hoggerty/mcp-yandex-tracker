// Package autoaction exposes Yandex Tracker queue auto-action operations as MCP tools.
package autoaction

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicautoaction "github.com/hoggerty/yandex-tracker-mcp/internal/logic/autoaction"
)

func Tools(svc *logicautoaction.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: logicautoaction.CreateSafety},
		{Tool: &getTool{svc: svc}, Safety: logicautoaction.GetSafety},
		{Tool: &logsListTool{svc: svc}, Safety: logicautoaction.LogsListSafety},
		{Tool: &logGetTool{svc: svc}, Safety: logicautoaction.LogGetSafety},
	}
}
