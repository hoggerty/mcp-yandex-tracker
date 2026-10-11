// Package project exposes Yandex Tracker project operations as MCP tools.
package project

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicproject "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/project"
)

func Tools(svc *logicproject.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &ListTool{Svc: svc}, Safety: svc.ListSafety()},
		{Tool: &QueuesListTool{Svc: svc}, Safety: svc.QueuesListSafety()},
		{Tool: &UpdateTool{Svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &DeleteTool{Svc: svc}, Safety: svc.DeleteSafety()},
	}
}
