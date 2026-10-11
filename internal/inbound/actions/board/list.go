package board

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicboard "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/board"
)

type ListTool struct{ Svc *logicboard.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("boards_list",
		mcp.WithDescription("List Yandex Tracker agile boards visible to the current user."),
		mcp.WithNumber("per_page", mcp.Description("Results per page, max 500")),
		mcp.WithNumber("id", mcp.Description("Last board ID from the previous page; omit for the first page")),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx, req.GetInt("per_page", 0), req.GetInt("id", 0)))
}
