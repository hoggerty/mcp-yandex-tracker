package trigger

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logictrigger "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/trigger"
)

type LogsListTool struct{ Svc *logictrigger.Service }

func (t *LogsListTool) Description() mcp.Tool {
	return mcp.NewTool("trigger_logs_list",
		mcp.WithDescription("View the webhook invocation log of a Yandex Tracker trigger."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("trigger_id", mcp.Required(), mcp.Description("Trigger id")),
		mcp.WithString("issue_id", mcp.Description("Filter by issue id")),
		mcp.WithNumber("limit", mcp.Description("Number of entries to return (default 10, max 100)")),
		mcp.WithString("from", mcp.Description("Start time, e.g. 2025-02-25T14:22:03.596+0000")),
		mcp.WithString("to", mcp.Description("End time, e.g. 2025-02-25T14:22:03.596+0000")),
	)
}

func (t *LogsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	triggerID, err := req.RequireString("trigger_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.LogsList(ctx, queueKey, triggerID,
		req.GetString("issue_id", ""),
		req.GetInt("limit", 0),
		req.GetString("from", ""),
		req.GetString("to", "")))
}
