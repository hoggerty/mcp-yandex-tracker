package board

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/board"
)

type EditTool struct{ Svc *logicboard.Service }

func (t *EditTool) Description() mcp.Tool {
	return mcp.NewTool("board_edit",
		mcp.WithDescription("Edit a Yandex Tracker agile board."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Board id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Board's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": \"...\", \"columns\": [...]}")),
	)
}

func (t *EditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	version, err := req.RequireInt("version")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Edit(ctx, id, version, fields))
}
