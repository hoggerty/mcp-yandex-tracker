// Package issuetype exposes Yandex Tracker issue type operations as MCP tools.
package issuetype

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicissuetype "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/issuetype"
)

func Tools(svc *logicissuetype.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
	}
}
