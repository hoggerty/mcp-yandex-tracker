// Package version exposes Yandex Tracker queue-version operations as MCP tools.
package version

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/version"
)

func Tools(svc *logicversion.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: logicversion.CreateSafety},
		{Tool: &queueListTool{svc: svc}, Safety: logicversion.QueueListSafety},
		{Tool: &getTool{svc: svc}, Safety: logicversion.GetSafety},
		{Tool: &editTool{svc: svc}, Safety: logicversion.EditSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicversion.DeleteSafety},
	}
}
