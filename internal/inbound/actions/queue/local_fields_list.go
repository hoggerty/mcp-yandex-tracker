package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicqueue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/queue"
)

type LocalFieldsListTool struct{ Svc *logicqueue.Service }

func (t *LocalFieldsListTool) Description() mcp.Tool {
	return mcp.NewTool("queue_local_fields_list",
		mcp.WithDescription("List a Yandex Tracker queue's local fields."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
	)
}

func (t *LocalFieldsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.LocalFieldsList(ctx, queueKey))
}
