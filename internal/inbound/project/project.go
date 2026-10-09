// Package project exposes Yandex Tracker project operations as MCP tools.
package project

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicproject "github.com/hoggerty/yandex-tracker-mcp/internal/logic/project"
)

func Tools(svc *logicproject.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &createTool{svc: svc}, Safety: logicproject.CreateSafety},
		{Tool: &getTool{svc: svc}, Safety: logicproject.GetSafety},
		{Tool: &listTool{svc: svc}, Safety: logicproject.ListSafety},
		{Tool: &queuesListTool{svc: svc}, Safety: logicproject.QueuesListSafety},
		{Tool: &updateTool{svc: svc}, Safety: logicproject.UpdateSafety},
		{Tool: &deleteTool{svc: svc}, Safety: logicproject.DeleteSafety},
	}
}
