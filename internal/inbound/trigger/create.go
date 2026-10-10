package trigger

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logictrigger "github.com/hoggerty/yandex-tracker-mcp/internal/logic/trigger"
)

type createTool struct{ svc *logictrigger.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("trigger_create",
		mcp.WithDescription("Create a new trigger on a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Trigger definition, e.g. {\"name\": \"...\", \"actions\": [...], \"conditions\": [...], \"active\": true}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, queueKey, fields))
}
