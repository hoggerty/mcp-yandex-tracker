// Package issuetype exposes Yandex Tracker issue type operations as MCP tools.
package issuetype

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicissuetype "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issuetype"
)

func Tools(svc *logicissuetype.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &listTool{svc: svc}, Safety: logicissuetype.ListSafety},
		{Tool: &createTool{svc: svc}, Safety: logicissuetype.CreateSafety},
		{Tool: &editTool{svc: svc}, Safety: logicissuetype.EditSafety},
	}
}
