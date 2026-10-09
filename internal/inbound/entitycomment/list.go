package entitycomment

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitycomment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitycomment"
)

type listTool struct{ svc *logicentitycomment.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("entitycomment_list",
		mcp.WithDescription("List comments on a Yandex Tracker entity."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("expand", mcp.Description("Optional: \"all\" or \"html\"")),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.List(ctx, entityType, id, req.GetString("expand", "")))
}
