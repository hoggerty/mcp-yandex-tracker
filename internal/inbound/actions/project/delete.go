package project

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicproject "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/project"
)

type DeleteTool struct{ Svc *logicproject.Service }

func (t *DeleteTool) Description() mcp.Tool {
	return mcp.NewTool("project_delete",
		mcp.WithDescription("Delete a Yandex Tracker project."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Project id")),
	)
}

func (t *DeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Delete(ctx, id))
}
