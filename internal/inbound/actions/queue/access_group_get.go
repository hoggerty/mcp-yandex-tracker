package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicqueue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/queue"
)

type AccessGroupGetTool struct{ Svc *logicqueue.Service }

func (t *AccessGroupGetTool) Description() mcp.Tool {
	return mcp.NewTool("queue_access_group_get",
		mcp.WithDescription("Get a group's permissions for a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("group_id", mcp.Required(), mcp.Description("Group id")),
	)
}

func (t *AccessGroupGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	groupID, err := req.RequireString("group_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.AccessGroupGet(ctx, queueKey, groupID))
}
