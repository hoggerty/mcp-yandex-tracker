package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/queue"
)

type localFieldsListTool struct{ svc *logicqueue.Service }

func (t *localFieldsListTool) Description() mcp.Tool {
	return mcp.NewTool("queue_local_fields_list",
		mcp.WithDescription("List a Yandex Tracker queue's local fields."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
	)
}

func (t *localFieldsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.LocalFieldsList(ctx, queueKey))
}
