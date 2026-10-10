// Package sprint exposes Yandex Tracker sprint operations as MCP tools.
package sprint

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicsprint "github.com/hoggerty/yandex-tracker-mcp/internal/logic/sprint"
)

func Tools(svc *logicsprint.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &startTool{svc: svc}, Safety: svc.StartSafety()},
		{Tool: &archiveTool{svc: svc}, Safety: svc.ArchiveSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
