package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicqueue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/queue"
)

type ListTool struct{ Svc *logicqueue.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("queues_list",
		mcp.WithDescription("List Yandex Tracker queues visible to the current user."),
		mcp.WithNumber("per_page", mcp.Description("Results per page")),
		mcp.WithNumber("page", mcp.Description("Page number, starting at 1")),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx, req.GetInt("per_page", 0), req.GetInt("page", 0)))
}
