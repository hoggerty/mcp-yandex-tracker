// Package sprint exposes Yandex Tracker sprint operations as MCP tools.
package sprint

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicsprint "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/sprint"
)

func Tools(svc *logicsprint.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &StartTool{Svc: svc}, Safety: svc.StartSafety()},
		{Tool: &ArchiveTool{Svc: svc}, Safety: svc.ArchiveSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
