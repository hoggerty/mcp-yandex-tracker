package trigger

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logictrigger "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/trigger"
)

type ListTool struct{ Svc *logictrigger.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("trigger_list",
		mcp.WithDescription("List the triggers configured on a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithNumber("per_page", mcp.Description("Results per page")),
		mcp.WithString("last_id", mcp.Description("Last trigger id from the previous page, for pagination")),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.List(ctx, queueKey, req.GetInt("per_page", 0), req.GetString("last_id", "")))
}
