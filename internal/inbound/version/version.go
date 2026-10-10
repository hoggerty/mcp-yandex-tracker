// Package version exposes Yandex Tracker queue-version operations as MCP tools.
package version

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/version"
)

func Tools(svc *logicversion.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &queueListTool{svc: svc}, Safety: svc.QueueListSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &editTool{svc: svc}, Safety: svc.EditSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
