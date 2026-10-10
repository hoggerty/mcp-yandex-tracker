package project

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicproject "github.com/hoggerty/yandex-tracker-mcp/internal/logic/project"
)

type deleteTool struct{ svc *logicproject.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("project_delete",
		mcp.WithDescription("Delete a Yandex Tracker project."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Project id")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, id))
}
