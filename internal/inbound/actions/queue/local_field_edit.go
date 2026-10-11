package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicqueue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/queue"
)

type LocalFieldEditTool struct{ Svc *logicqueue.Service }

func (t *LocalFieldEditTool) Description() mcp.Tool {
	return mcp.NewTool("queue_local_field_edit",
		mcp.WithDescription("Edit a queue-local issue field."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("field_key", mcp.Required(), mcp.Description("Local field key")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": {\"en\": \"...\", \"ru\": \"...\"}, \"order\": 102, \"readonly\": true}")),
	)
}

func (t *LocalFieldEditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fieldKey, err := req.RequireString("field_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.LocalFieldEdit(ctx, queueKey, fieldKey, fields))
}
