package column

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccolumn "github.com/hoggerty/yandex-tracker-mcp/internal/logic/column"
)

type getTool struct{ svc *logiccolumn.Service }

func (t *getTool) Description() mcp.Tool {
	return mcp.NewTool("column_get",
		mcp.WithDescription("Get a single column of a Yandex Tracker agile board."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("board_id", mcp.Required(), mcp.Description("Board id")),
		mcp.WithString("column_id", mcp.Required(), mcp.Description("Column id")),
	)
}

func (t *getTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	boardID, err := req.RequireString("board_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	columnID, err := req.RequireString("column_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Get(ctx, boardID, columnID))
}
