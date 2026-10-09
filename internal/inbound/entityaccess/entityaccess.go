// Package entityaccess exposes Yandex Tracker entity access/permissions operations as MCP tools.
package entityaccess

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicentityaccess "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityaccess"
)

func Tools(svc *logicentityaccess.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &accessGetTool{svc: svc}, Safety: logicentityaccess.AccessGetSafety},
		{Tool: &permissionsGetTool{svc: svc}, Safety: logicentityaccess.PermissionsGetSafety},
		{Tool: &accessEditTool{svc: svc}, Safety: logicentityaccess.AccessEditSafety},
		{Tool: &permissionsEditTool{svc: svc}, Safety: logicentityaccess.PermissionsEditSafety},
	}
}
