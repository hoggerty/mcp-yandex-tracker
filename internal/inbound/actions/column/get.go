package column

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logiccolumn "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/column"
)

type GetTool struct{ Svc *logiccolumn.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("column_get",
		mcp.WithDescription("Get a single column of a Yandex Tracker agile board."),
		mcp.WithString("board_id", mcp.Required(), mcp.Description("Board id")),
		mcp.WithString("column_id", mcp.Required(), mcp.Description("Column id")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	boardID, err := req.RequireString("board_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	columnID, err := req.RequireString("column_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, boardID, columnID))
}
