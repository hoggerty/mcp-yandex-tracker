package column

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logiccolumn "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/column"
)

type ListTool struct{ Svc *logiccolumn.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("column_list",
		mcp.WithDescription("List the columns of a Yandex Tracker agile board."),
		mcp.WithString("board_id", mcp.Required(), mcp.Description("Board id")),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	boardID, err := req.RequireString("board_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.List(ctx, boardID))
}
