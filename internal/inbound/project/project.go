// Package project exposes Yandex Tracker project operations as MCP tools.
package project

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicproject "github.com/hoggerty/yandex-tracker-mcp/internal/logic/project"
)

func Tools(svc *logicproject.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &listTool{svc: svc}, Safety: svc.ListSafety()},
		{Tool: &queuesListTool{svc: svc}, Safety: svc.QueuesListSafety()},
		{Tool: &updateTool{svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &deleteTool{svc: svc}, Safety: svc.DeleteSafety()},
	}
}
