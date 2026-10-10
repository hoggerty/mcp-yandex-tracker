package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/queue"
)

type FieldsListTool struct{ Svc *logicqueue.Service }

func (t *FieldsListTool) Description() mcp.Tool {
	return mcp.NewTool("queue_fields_list",
		mcp.WithDescription("Get the required fields configured for a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
	)
}

func (t *FieldsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.FieldsList(ctx, queueKey))
}
