package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/queue"
)

type LocalFieldGetTool struct{ Svc *logicqueue.Service }

func (t *LocalFieldGetTool) Description() mcp.Tool {
	return mcp.NewTool("queue_local_field_get",
		mcp.WithDescription("Get a single local field of a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("field_key", mcp.Required(), mcp.Description("Local field key")),
	)
}

func (t *LocalFieldGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fieldKey, err := req.RequireString("field_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.LocalFieldGet(ctx, queueKey, fieldKey))
}
