package board

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicboard "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/board"
)

type CreateTool struct{ Svc *logicboard.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("board_create",
		mcp.WithDescription("Create a new Yandex Tracker agile board."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Board definition, e.g. {\"name\": \"...\", \"columns\": [...], \"queues\": [\"TEST\"]}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, fields))
}
