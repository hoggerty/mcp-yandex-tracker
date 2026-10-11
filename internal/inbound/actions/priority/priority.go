// Package priority exposes Yandex Tracker issue priority operations as MCP tools.
package priority

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicpriority "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/priority"
)

func Tools(svc *logicpriority.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
	}
}
