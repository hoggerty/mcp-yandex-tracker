package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/queue"
)

type LocalFieldCreateTool struct{ Svc *logicqueue.Service }

func (t *LocalFieldCreateTool) Description() mcp.Tool {
	return mcp.NewTool("queue_local_field_create",
		mcp.WithDescription("Create a queue-local issue field."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Field definition, e.g. {\"id\": \"myField\", \"name\": {\"en\": \"My Field\", \"ru\": \"Моё поле\"}, \"category\": \"<category id>\", \"type\": \"ru.yandex.startrek.core.fields.StringFieldType\"}")),
	)
}

func (t *LocalFieldCreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.LocalFieldCreate(ctx, queueKey, fields))
}
