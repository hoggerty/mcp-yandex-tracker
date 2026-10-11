// Package report exposes Yandex Tracker issue-report operations as MCP tools.
package report

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicreport "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/report"
)

func Tools(svc *logicreport.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &SearchTool{Svc: svc}, Safety: svc.SearchSafety()},
	}
}
