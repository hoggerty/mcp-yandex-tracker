package board

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/board"
)

type getTool struct{ svc *logicboard.Service }

func (t *getTool) Description() mcp.Tool {
	return mcp.NewTool("board_get",
		mcp.WithDescription("Get a Yandex Tracker agile board by id."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Board id")),
	)
}

func (t *getTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Get(ctx, id))
}
