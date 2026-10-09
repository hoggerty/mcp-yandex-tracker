package board

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/board"
)

type listTool struct{ svc *logicboard.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("boards_list",
		mcp.WithDescription("List Yandex Tracker agile boards visible to the current user."),
	)
}

func (t *listTool) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx))
}
