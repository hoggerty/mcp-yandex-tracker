package entitychecklist

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitychecklist "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitychecklist"
)

type itemDeleteTool struct{ svc *logicentitychecklist.Service }

func (t *itemDeleteTool) Description() mcp.Tool {
	return mcp.NewTool("entitychecklist_item_delete",
		mcp.WithDescription("Delete a single checklist item from a Yandex Tracker entity."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("item_id", mcp.Required(), mcp.Description("Checklist item id")),
	)
}

func (t *itemDeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	return result.JSON(t.svc.ItemDelete(ctx, entityType, id, itemID))
}
