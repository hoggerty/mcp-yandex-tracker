package entitycomment

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitycomment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitycomment"
)

type deleteTool struct{ svc *logicentitycomment.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("entitycomment_delete",
		mcp.WithDescription("Delete a comment on a Yandex Tracker entity."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("comment_id", mcp.Required(), mcp.Description("Comment id")),
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
	commentID, err := req.RequireString("comment_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, entityType, id, commentID))
}
