package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/queue"
)

type GetTool struct{ Svc *logicqueue.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("queue_get",
		mcp.WithDescription("Retrieve a Yandex Tracker queue by its key."),
		mcp.WithString("queue_key", mcp.Required(), mcp.Description("Queue key, e.g. QUEUE")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, queueKey))
}
