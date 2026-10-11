package board

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicboard "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/board"
)

type GetTool struct{ Svc *logicboard.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("board_get",
		mcp.WithDescription("Get a Yandex Tracker agile board by id."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Board id")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, id))
}
