package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/queue"
)

type AccessUserGetTool struct{ Svc *logicqueue.Service }

func (t *AccessUserGetTool) Description() mcp.Tool {
	return mcp.NewTool("queue_access_user_get",
		mcp.WithDescription("Get a user's permissions for a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("user_id", mcp.Required(), mcp.Description("User id or login")),
	)
}

func (t *AccessUserGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	userID, err := req.RequireString("user_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.AccessUserGet(ctx, queueKey, userID))
}
