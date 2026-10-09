package column

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccolumn "github.com/hoggerty/yandex-tracker-mcp/internal/logic/column"
)

type editTool struct{ svc *logiccolumn.Service }

func (t *editTool) Description() mcp.Tool {
	return mcp.NewTool("column_edit",
		mcp.WithDescription("Edit a column of a Yandex Tracker agile board."),
		mcp.WithString("board_id", mcp.Required(), mcp.Description("Board id")),
		mcp.WithString("column_id", mcp.Required(), mcp.Description("Column id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Board's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": \"...\", \"statuses\": [\"open\"]}")),
	)
}

func (t *editTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	boardID, err := req.RequireString("board_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	columnID, err := req.RequireString("column_id")
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
	return result.JSON(t.svc.Edit(ctx, boardID, columnID, version, fields))
}
