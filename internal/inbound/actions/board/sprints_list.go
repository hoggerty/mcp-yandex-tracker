package board

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/board"
)

type SprintsListTool struct{ Svc *logicboard.Service }

func (t *SprintsListTool) Description() mcp.Tool {
	return mcp.NewTool("board_sprints_list",
		mcp.WithDescription("List sprints that belong to a Yandex Tracker agile board."),
		mcp.WithString("board_id", mcp.Required(), mcp.Description("Board ID")),
	)
}

func (t *SprintsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	boardID, err := req.RequireString("board_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.ListSprints(ctx, boardID))
}
