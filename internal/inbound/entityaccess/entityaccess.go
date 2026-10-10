// Package entityaccess exposes Yandex Tracker entity access/permissions operations as MCP tools.
package entityaccess

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentityaccess "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityaccess"
)

func Tools(svc *logicentityaccess.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &accessGetTool{svc: svc}, Safety: svc.AccessGetSafety()},
		{Tool: &permissionsGetTool{svc: svc}, Safety: svc.PermissionsGetSafety()},
		{Tool: &accessEditTool{svc: svc}, Safety: svc.AccessEditSafety()},
		{Tool: &permissionsEditTool{svc: svc}, Safety: svc.PermissionsEditSafety()},
	}
}
