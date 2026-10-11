package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logiccomponent "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/component"
)

type QueueListTool struct{ Svc *logiccomponent.Service }

func (t *QueueListTool) Description() mcp.Tool {
	return mcp.NewTool("component_queue_list",
		mcp.WithDescription("List the components of a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
	)
}

func (t *QueueListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.QueueList(ctx, queueKey))
}
