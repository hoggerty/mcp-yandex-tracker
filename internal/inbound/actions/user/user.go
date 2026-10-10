// Package user exposes Yandex Tracker user operations as MCP tools.
package user

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicuser "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/user"
)

func Tools(svc *logicuser.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CurrentTool{Svc: svc}, Safety: svc.CurrentSafety()},
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &ListRelativeTool{Svc: svc}, Safety: svc.ListRelativeSafety()},
	}
}
