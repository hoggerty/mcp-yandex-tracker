package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/queue"
)

type tagsListTool struct{ svc *logicqueue.Service }

func (t *tagsListTool) Description() mcp.Tool {
	return mcp.NewTool("queue_tags_list",
		mcp.WithDescription("Get the tags defined on a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
	)
}

func (t *tagsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.TagsList(ctx, queueKey))
}
