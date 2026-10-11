package workflow

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicworkflow "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/workflow"
)

type QueueListTool struct{ Svc *logicworkflow.Service }

func (t *QueueListTool) Description() mcp.Tool {
	return mcp.NewTool("workflow_queue_list",
		mcp.WithDescription("Get the workflow assigned to each issue type in a queue."),
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
