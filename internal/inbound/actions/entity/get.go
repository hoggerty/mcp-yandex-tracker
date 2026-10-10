package entity

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entity"
)

type GetTool struct{ Svc *logicentity.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("entity_get",
		mcp.WithDescription("Get a Yandex Tracker entity (project, portfolio, or goal) by id."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("expand", mcp.Description("Optional: \"attachments\"")),
		mcp.WithString("fields", mcp.Description("Optional comma-separated field names to include")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, entityType, id, req.GetString("expand", ""), req.GetString("fields", "")))
}
