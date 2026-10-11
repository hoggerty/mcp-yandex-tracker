package trigger

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logictrigger "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/trigger"
)

type GetTool struct{ Svc *logictrigger.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("trigger_get",
		mcp.WithDescription("Get the parameters of a single Yandex Tracker trigger."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("trigger_id", mcp.Required(), mcp.Description("Trigger id")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	triggerID, err := req.RequireString("trigger_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, queueKey, triggerID))
}
