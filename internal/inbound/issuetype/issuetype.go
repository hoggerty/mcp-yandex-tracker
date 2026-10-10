// Package issuetype exposes Yandex Tracker issue type operations as MCP tools.
package issuetype

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicissuetype "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issuetype"
)

func Tools(svc *logicissuetype.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
	}
}
