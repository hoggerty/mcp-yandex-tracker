package entitycomment

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitycomment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entitycomment"
)

type ListTool struct{ Svc *logicentitycomment.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("entitycomment_list",
		mcp.WithDescription("List comments on a Yandex Tracker entity."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("expand", mcp.Description("Optional: \"all\" or \"html\"")),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.List(ctx, entityType, id, req.GetString("expand", "")))
}
