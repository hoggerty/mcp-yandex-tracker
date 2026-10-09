package entitychecklist

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitychecklist "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitychecklist"
)

type deleteTool struct{ svc *logicentitychecklist.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("entitychecklist_delete",
		mcp.WithDescription("Delete all checklist items from a Yandex Tracker entity."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, entityType, id))
}
