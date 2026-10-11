package column

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logiccolumn "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/column"
)

type CreateTool struct{ Svc *logiccolumn.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("column_create",
		mcp.WithDescription("Create a new column on a Yandex Tracker agile board."),
		mcp.WithString("board_id", mcp.Required(), mcp.Description("Board id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Board's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Column definition, e.g. {\"name\": \"...\", \"statuses\": [\"open\"]}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	boardID, err := req.RequireString("board_id")
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
	return result.JSON(t.Svc.Create(ctx, boardID, version, fields))
}
