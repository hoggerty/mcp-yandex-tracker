// Package report exposes Yandex Tracker issue-report operations as MCP tools.
package report

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicreport "github.com/hoggerty/yandex-tracker-mcp/internal/logic/report"
)

func Tools(svc *logicreport.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: logicreport.CreateSafety},
		{Tool: &searchTool{svc: svc}, Safety: logicreport.SearchSafety},
	}
}
