package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/queue"
)

type listTool struct{ svc *logicqueue.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("queues_list",
		mcp.WithDescription("List Yandex Tracker queues visible to the current user."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithNumber("per_page", mcp.Description("Results per page")),
		mcp.WithNumber("page", mcp.Description("Page number, starting at 1")),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx, req.GetInt("per_page", 0), req.GetInt("page", 0)))
}
