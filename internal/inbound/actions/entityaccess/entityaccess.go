// Package entityaccess exposes Yandex Tracker entity access/permissions operations as MCP tools.
package entityaccess

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicentityaccess "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/entityaccess"
)

func Tools(svc *logicentityaccess.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &AccessGetTool{Svc: svc}, Safety: svc.AccessGetSafety()},
		{Tool: &PermissionsGetTool{Svc: svc}, Safety: svc.PermissionsGetSafety()},
		{Tool: &AccessEditTool{Svc: svc}, Safety: svc.AccessEditSafety()},
		{Tool: &PermissionsEditTool{Svc: svc}, Safety: svc.PermissionsEditSafety()},
	}
}
