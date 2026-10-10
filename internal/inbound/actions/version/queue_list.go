package version

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/version"
)

type QueueListTool struct{ Svc *logicversion.Service }

func (t *QueueListTool) Description() mcp.Tool {
	return mcp.NewTool("version_queue_list",
		mcp.WithDescription("List the versions of a Yandex Tracker queue."),
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
