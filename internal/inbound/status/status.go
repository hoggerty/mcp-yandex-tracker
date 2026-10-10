// Package status exposes Yandex Tracker issue status operations as MCP tools.
package status

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicstatus "github.com/hoggerty/yandex-tracker-mcp/internal/logic/status"
)

func Tools(svc *logicstatus.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
	}
}
