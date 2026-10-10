// Package user exposes Yandex Tracker user operations as MCP tools.
package user

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicuser "github.com/hoggerty/yandex-tracker-mcp/internal/logic/user"
)

func Tools(svc *logicuser.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &currentTool{svc: svc}, Safety: svc.CurrentSafety()},
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &listRelativeTool{svc: svc}, Safety: svc.ListRelativeSafety()},
	}
}
