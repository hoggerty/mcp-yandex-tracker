package entitychecklist

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitychecklist "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitychecklist"
)

type editTool struct{ svc *logicentitychecklist.Service }

func (t *editTool) Description() mcp.Tool {
	return mcp.NewTool("entitychecklist_edit",
		mcp.WithDescription("Replace the full checklist on a Yandex Tracker entity with the given items (bulk edit)."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithArray("items", mcp.Required(), mcp.Description("Array of checklist item objects, e.g. [{\"text\": \"...\", \"checked\": false}]")),
	)
}

func (t *editTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	raw, ok := req.GetArguments()["items"].([]any)
	if !ok || len(raw) == 0 {
		return mcp.NewToolResultError("items must be a non-empty array"), nil
	}
	items := make([]map[string]any, 0, len(raw))
	for _, v := range raw {
		item, ok := v.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("each item must be an object"), nil
		}
		items = append(items, item)
	}
	return result.JSON(t.svc.Edit(ctx, entityType, id, items))
}
