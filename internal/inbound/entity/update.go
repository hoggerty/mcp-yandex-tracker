package entity

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entity"
)

type updateTool struct{ svc *logicentity.Service }

func (t *updateTool) Description() mcp.Tool {
	return mcp.NewTool("entity_update",
		mcp.WithDescription("Edit a Yandex Tracker entity (project, portfolio, or goal)."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithObject("body", mcp.Required(), mcp.Description("e.g. {\"fields\": {\"summary\": \"...\", \"description\": \"...\"}, \"comment\": \"...\", \"links\": [...]}")),
	)
}

func (t *updateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	body, ok := req.GetArguments()["body"].(map[string]any)
	if !ok || len(body) == 0 {
		return mcp.NewToolResultError("body must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Update(ctx, entityType, id, body))
}
