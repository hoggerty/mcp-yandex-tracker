package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicqueue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/queue"
)

type createTool struct{ svc *logicqueue.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("queue_create",
		mcp.WithDescription("Create a new Yandex Tracker queue."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Queue definition, e.g. {\"key\": \"DESIGN\", \"name\": \"Design\", \"lead\": \"username\", \"defaultType\": \"task\", \"defaultPriority\": \"normal\", \"issueTypesConfig\": [{\"issueType\": \"task\", \"workflow\": \"hrPresetWorkflow\", \"resolutions\": [\"fixed\"]}]}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
