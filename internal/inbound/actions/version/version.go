// Package version exposes Yandex Tracker queue-version operations as MCP tools.
package version

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/version"
)

func Tools(svc *logicversion.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &QueueListTool{Svc: svc}, Safety: svc.QueueListSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &EditTool{Svc: svc}, Safety: svc.EditSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
