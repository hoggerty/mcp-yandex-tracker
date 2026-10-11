package entitychecklist

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicentitychecklist "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/entitychecklist"
)

type ItemDeleteTool struct{ Svc *logicentitychecklist.Service }

func (t *ItemDeleteTool) Description() mcp.Tool {
	return mcp.NewTool("entitychecklist_item_delete",
		mcp.WithDescription("Delete a single checklist item from a Yandex Tracker entity."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("item_id", mcp.Required(), mcp.Description("Checklist item id")),
	)
}

func (t *ItemDeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	itemID, err := req.RequireString("item_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.ItemDelete(ctx, entityType, id, itemID))
}
