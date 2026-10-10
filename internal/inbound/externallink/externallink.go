// Package externallink exposes Yandex Tracker external-link operations as MCP tools.
package externallink

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicexternallink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/externallink"
)

func Tools(svc *logicexternallink.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &applicationsListTool{svc: svc}, Safety: svc.ApplicationsListSafety()},
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &addTool{svc: svc}, Safety: svc.AddSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
