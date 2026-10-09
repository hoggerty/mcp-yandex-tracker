package entitylink

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitylink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitylink"
)

type deleteTool struct{ svc *logicentitylink.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("entitylink_delete",
		mcp.WithDescription("Delete a link from a Yandex Tracker entity."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("right_id", mcp.Required(), mcp.Description("Id of the linked entity/issue on the right side of the relationship")),
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
	rightID, err := req.RequireString("right_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, entityType, id, rightID))
}
