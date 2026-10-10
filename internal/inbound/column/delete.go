package column

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccolumn "github.com/hoggerty/yandex-tracker-mcp/internal/logic/column"
)

type deleteTool struct{ svc *logiccolumn.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("column_delete",
		mcp.WithDescription("Delete a column of a Yandex Tracker agile board."),
		mcp.WithString("board_id", mcp.Required(), mcp.Description("Board id")),
		mcp.WithString("column_id", mcp.Required(), mcp.Description("Column id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Board's current version (for optimistic locking)")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	return result.JSON(t.svc.Delete(ctx, boardID, columnID, version))
}
