package entitychecklist

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitychecklist "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitychecklist"
)

type addTool struct{ svc *logicentitychecklist.Service }

func (t *addTool) Description() mcp.Tool {
	return mcp.NewTool("entitychecklist_add",
		mcp.WithDescription("Add a checklist item to a Yandex Tracker entity."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"text\": \"...\", \"checked\": false, \"assignee\": \"login\"}")),
	)
}

func (t *addTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Add(ctx, entityType, id, fields))
}
