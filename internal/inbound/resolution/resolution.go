// Package resolution exposes Yandex Tracker issue resolution operations as MCP tools.
package resolution

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicresolution "github.com/hoggerty/yandex-tracker-mcp/internal/logic/resolution"
)

func Tools(svc *logicresolution.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: logicresolution.ListSafety},
		{Tool: &createTool{svc: svc}, Safety: logicresolution.CreateSafety},
		{Tool: &editTool{svc: svc}, Safety: logicresolution.EditSafety},
	}
}
