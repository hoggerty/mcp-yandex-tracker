package entity

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entity"
)

type BulkUpdateTool struct{ Svc *logicentity.Service }

func (t *BulkUpdateTool) Description() mcp.Tool {
	return mcp.NewTool("entity_bulk_update",
		mcp.WithDescription("Bulk-edit fields, comment, or links on multiple Yandex Tracker entities of the same type."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithObject("body", mcp.Required(), mcp.Description("e.g. {\"metaEntities\": [\"ID1\", \"ID2\"], \"values\": {\"fields\": {\"FIELD_KEY\": \"VALUE\"}, \"comment\": \"...\", \"links\": [...]}}")),
	)
}

func (t *BulkUpdateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	body, ok := req.GetArguments()["body"].(map[string]any)
	if !ok || len(body) == 0 {
		return mcp.NewToolResultError("body must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.BulkUpdate(ctx, entityType, body))
}
