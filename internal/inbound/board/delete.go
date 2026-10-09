package board

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/board"
)

type deleteTool struct{ svc *logicboard.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("board_delete",
		mcp.WithDescription("Delete a Yandex Tracker agile board."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("id", mcp.Required(), mcp.Description("Board id")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, id))
}
