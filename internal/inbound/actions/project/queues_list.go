package project

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicproject "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/project"
)

type QueuesListTool struct{ Svc *logicproject.Service }

func (t *QueuesListTool) Description() mcp.Tool {
	return mcp.NewTool("project_queues_list",
		mcp.WithDescription("List the queues attached to a Yandex Tracker project."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Project id")),
		mcp.WithString("expand", mcp.Description("Optional: all, projects, components, versions, types, team, workflows, fields")),
	)
}

func (t *QueuesListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.QueuesList(ctx, id, req.GetString("expand", "")))
}
