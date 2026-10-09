package autoaction

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicautoaction "github.com/hoggerty/yandex-tracker-mcp/internal/logic/autoaction"
)

type createTool struct{ svc *logicautoaction.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("autoaction_create",
		mcp.WithDescription("Create a new auto action on a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Auto action definition, e.g. {\"name\": \"...\", \"filter\": {...}, \"actions\": [...], \"active\": true, \"intervalMillis\": 3600000}")),
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
