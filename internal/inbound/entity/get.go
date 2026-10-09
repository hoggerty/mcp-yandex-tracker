package entity

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entity"
)

type getTool struct{ svc *logicentity.Service }

func (t *getTool) Description() mcp.Tool {
	return mcp.NewTool("entity_get",
		mcp.WithDescription("Get a Yandex Tracker entity (project, portfolio, or goal) by id."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("expand", mcp.Description("Optional: \"attachments\"")),
		mcp.WithString("fields", mcp.Description("Optional comma-separated field names to include")),
	)
}

func (t *getTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Get(ctx, entityType, id, req.GetString("expand", ""), req.GetString("fields", "")))
}
