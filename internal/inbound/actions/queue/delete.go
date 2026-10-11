package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicqueue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/queue"
)

type DeleteTool struct{ Svc *logicqueue.Service }

func (t *DeleteTool) Description() mcp.Tool {
	return mcp.NewTool("queue_delete",
		mcp.WithDescription("Delete a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
	)
}

func (t *DeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Delete(ctx, queueKey))
}
