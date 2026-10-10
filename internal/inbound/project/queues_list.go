package project

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicproject "github.com/hoggerty/yandex-tracker-mcp/internal/logic/project"
)

type queuesListTool struct{ svc *logicproject.Service }

func (t *queuesListTool) Description() mcp.Tool {
	return mcp.NewTool("project_queues_list",
		mcp.WithDescription("List the queues attached to a Yandex Tracker project."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Project id")),
		mcp.WithString("expand", mcp.Description("Optional: all, projects, components, versions, types, team, workflows, fields")),
	)
}

func (t *queuesListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.QueuesList(ctx, id, req.GetString("expand", "")))
}
