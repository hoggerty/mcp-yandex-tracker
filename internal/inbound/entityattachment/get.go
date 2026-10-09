package entityattachment

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentityattachment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityattachment"
)

type getTool struct {
	svc *logicentityattachment.Service
}

func (t *getTool) Description() mcp.Tool {
	return mcp.NewTool("entityattachment_get",
		mcp.WithDescription("Get a single attachment's metadata on a Yandex Tracker entity."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("file_id", mcp.Required(), mcp.Description("Attachment file id")),
	)
}

func (t *getTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fileID, err := req.RequireString("file_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Get(ctx, entityType, id, fileID))
}
