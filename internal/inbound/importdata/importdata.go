// Package importdata exposes Yandex Tracker historical-data import operations as MCP tools.
package importdata

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicimportdata "github.com/hoggerty/yandex-tracker-mcp/internal/logic/importdata"
)

func Tools(svc *logicimportdata.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ticketTool{svc: svc}, Safety: logicimportdata.TicketSafety},
		{Tool: &commentTool{svc: svc}, Safety: logicimportdata.CommentSafety},
		{Tool: &linkTool{svc: svc}, Safety: logicimportdata.LinkSafety},
		{Tool: &worklogTool{svc: svc}, Safety: logicimportdata.WorklogSafety},
	}
}
