package entitychecklist

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicentitychecklist "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/entitychecklist"
)

type ItemMoveTool struct{ Svc *logicentitychecklist.Service }

func (t *ItemMoveTool) Description() mcp.Tool {
	return mcp.NewTool("entitychecklist_item_move",
		mcp.WithDescription("Move a checklist item to a new position on a Yandex Tracker entity."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("item_id", mcp.Required(), mcp.Description("Checklist item id to move")),
		mcp.WithString("before_id", mcp.Required(), mcp.Description("Checklist item id to place the moved item before")),
	)
}

func (t *ItemMoveTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	beforeID, err := req.RequireString("before_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.ItemMove(ctx, entityType, id, itemID, beforeID))
}
