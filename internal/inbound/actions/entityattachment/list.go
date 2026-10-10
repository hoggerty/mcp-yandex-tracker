package entityattachment

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentityattachment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entityattachment"
)

type ListTool struct {
	Svc *logicentityattachment.Service
}

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("entityattachment_list",
		mcp.WithDescription("List attachments on a Yandex Tracker entity (metadata only)."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
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
	return result.JSON(t.Svc.List(ctx, entityType, id))
}
