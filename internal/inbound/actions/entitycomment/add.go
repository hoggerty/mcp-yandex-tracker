package entitycomment

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitycomment "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entitycomment"
)

type AddTool struct{ Svc *logicentitycomment.Service }

func (t *AddTool) Description() mcp.Tool {
	return mcp.NewTool("entitycomment_add",
		mcp.WithDescription("Add a comment to a Yandex Tracker entity (project, portfolio, or goal)."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"text\": \"...\", \"summonees\": [\"login\"]}")),
	)
}

func (t *AddTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	return result.JSON(t.Svc.Add(ctx, entityType, id, fields))
}
