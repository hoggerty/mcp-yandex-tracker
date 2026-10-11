// Package resolution exposes Yandex Tracker issue resolution operations as MCP tools.
package resolution

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicresolution "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/resolution"
)

func Tools(svc *logicresolution.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
	}
}
