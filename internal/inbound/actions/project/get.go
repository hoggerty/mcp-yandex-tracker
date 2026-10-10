package project

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicproject "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/project"
)

type GetTool struct{ Svc *logicproject.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("project_get",
		mcp.WithDescription("Get a Yandex Tracker project by id."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Project id")),
		mcp.WithString("expand", mcp.Description("Optional: \"queues\"")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, id, req.GetString("expand", "")))
}
