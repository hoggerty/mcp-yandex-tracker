package macro

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicmacro "github.com/hoggerty/yandex-tracker-mcp/internal/logic/macro"
)

type createTool struct{ svc *logicmacro.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("macro_create",
		mcp.WithDescription("Create a new macro on a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Macro definition, e.g. {\"name\": \"...\", \"body\": \"comment text\", \"issueUpdate\": {\"tags\": {\"add\": \"New tag\"}}}")),
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
