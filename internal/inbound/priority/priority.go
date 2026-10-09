// Package priority exposes Yandex Tracker issue priority operations as MCP tools.
package priority

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicpriority "github.com/hoggerty/yandex-tracker-mcp/internal/logic/priority"
)

func Tools(svc *logicpriority.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: logicpriority.ListSafety},
		{Tool: &createTool{svc: svc}, Safety: logicpriority.CreateSafety},
		{Tool: &editTool{svc: svc}, Safety: logicpriority.EditSafety},
	}
}
