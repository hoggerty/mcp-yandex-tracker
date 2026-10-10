package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/queue"
)

type RestoreTool struct{ Svc *logicqueue.Service }

func (t *RestoreTool) Description() mcp.Tool {
	return mcp.NewTool("queue_restore",
		mcp.WithDescription("Restore a previously deleted Yandex Tracker queue. Administrator-only."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
	)
}

func (t *RestoreTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Restore(ctx, queueKey))
}
