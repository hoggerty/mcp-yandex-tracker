package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/queue"
)

type AccessManageTool struct{ Svc *logicqueue.Service }

func (t *AccessManageTool) Description() mcp.Tool {
	return mcp.NewTool("queue_access_manage",
		mcp.WithDescription("Grant or revoke queue-level permissions (create/write/read/grant) for users, groups, or roles."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithObject("permissions", mcp.Required(), mcp.Description("Permission changes, e.g. {\"write\": {\"users\": {\"add\": [\"user1\"], \"remove\": [\"user2\"]}}}")),
	)
}

func (t *AccessManageTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	permissions, ok := req.GetArguments()["permissions"].(map[string]any)
	if !ok || len(permissions) == 0 {
		return mcp.NewToolResultError("permissions must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.AccessManage(ctx, queueKey, permissions))
}
