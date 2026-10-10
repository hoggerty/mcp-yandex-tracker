package trigger

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logictrigger "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/trigger"
)

type EditTool struct{ Svc *logictrigger.Service }

func (t *EditTool) Description() mcp.Tool {
	return mcp.NewTool("trigger_edit",
		mcp.WithDescription("Edit a Yandex Tracker trigger."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("trigger_id", mcp.Required(), mcp.Description("Trigger id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Trigger's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": \"...\", \"actions\": [...], \"conditions\": [...], \"active\": true}")),
	)
}

func (t *EditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	triggerID, err := req.RequireString("trigger_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	version, err := req.RequireInt("version")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Edit(ctx, queueKey, triggerID, version, fields))
}
