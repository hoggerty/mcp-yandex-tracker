// Package sprint exposes Yandex Tracker sprint operations as MCP tools.
package sprint

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicsprint "github.com/hoggerty/yandex-tracker-mcp/internal/logic/sprint"
)

func Tools(svc *logicsprint.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &getTool{svc: svc}, Safety: logicsprint.GetSafety},
		{Tool: &createTool{svc: svc}, Safety: logicsprint.CreateSafety},
		{Tool: &editTool{svc: svc}, Safety: logicsprint.EditSafety},
		{Tool: &startTool{svc: svc}, Safety: logicsprint.StartSafety},
		{Tool: &archiveTool{svc: svc}, Safety: logicsprint.ArchiveSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicsprint.DeleteSafety},
	}
}
