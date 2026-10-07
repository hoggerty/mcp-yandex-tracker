package board

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/result"
	"github.com/hoggerty/yandex-tracker-mcp/internal/trackerapi"
)

type sprintsListTool struct{ client *trackerapi.Client }

func (t *sprintsListTool) Description() mcp.Tool {
	return mcp.NewTool("board_sprints_list",
		mcp.WithDescription("List sprints that belong to a Yandex Tracker agile board."),
		mcp.WithString("board_id", mcp.Required(), mcp.Description("Board ID")),
	)
}

func (t *sprintsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	boardID, err := req.RequireString("board_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.client.ListBoardSprints(ctx, boardID))
}
