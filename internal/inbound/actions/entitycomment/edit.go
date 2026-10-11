package entitycomment

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicentitycomment "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/entitycomment"
)

type EditTool struct{ Svc *logicentitycomment.Service }

func (t *EditTool) Description() mcp.Tool {
	return mcp.NewTool("entitycomment_edit",
		mcp.WithDescription("Edit a comment on a Yandex Tracker entity."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithString("comment_id", mcp.Required(), mcp.Description("Comment id")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"text\": \"...\", \"summonees\": [\"login\"]}")),
	)
}

func (t *EditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Edit(ctx, entityType, id, commentID, fields))
}
