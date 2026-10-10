// Package status exposes Yandex Tracker issue status operations as MCP tools.
package status

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicstatus "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/status"
)

func Tools(svc *logicstatus.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
	}
}
