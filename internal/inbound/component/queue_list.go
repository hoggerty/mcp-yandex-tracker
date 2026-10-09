package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
)

type queueListTool struct{ svc *logiccomponent.Service }

func (t *queueListTool) Description() mcp.Tool {
	return mcp.NewTool("component_queue_list",
		mcp.WithDescription("List the components of a Yandex Tracker queue."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
	)
}

func (t *queueListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.QueueList(ctx, queueKey))
}
