// Package externallink exposes Yandex Tracker external-link operations as MCP tools.
package externallink

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicexternallink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/externallink"
)

func Tools(svc *logicexternallink.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &ApplicationsListTool{Svc: svc}, Safety: svc.ApplicationsListSafety()},
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &AddTool{Svc: svc}, Safety: svc.AddSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
