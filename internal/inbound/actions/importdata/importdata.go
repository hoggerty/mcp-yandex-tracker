// Package importdata exposes Yandex Tracker historical-data import operations as MCP tools.
package importdata

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicimportdata "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/importdata"
)

func Tools(svc *logicimportdata.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &TicketTool{Svc: svc}, Safety: svc.TicketSafety()},
		{Tool: &CommentTool{Svc: svc}, Safety: svc.CommentSafety()},
		{Tool: &LinkTool{Svc: svc}, Safety: svc.LinkSafety()},
		{Tool: &WorklogTool{Svc: svc}, Safety: svc.WorklogSafety()},
	}
}
