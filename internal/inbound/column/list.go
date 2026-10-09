package column

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccolumn "github.com/hoggerty/yandex-tracker-mcp/internal/logic/column"
)

type listTool struct{ svc *logiccolumn.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("column_list",
		mcp.WithDescription("List the columns of a Yandex Tracker agile board."),
		mcp.WithString("board_id", mcp.Required(), mcp.Description("Board id")),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	boardID, err := req.RequireString("board_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.List(ctx, boardID))
}
