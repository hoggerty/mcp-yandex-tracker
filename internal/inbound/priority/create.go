package priority

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicpriority "github.com/hoggerty/yandex-tracker-mcp/internal/logic/priority"
)

type createTool struct{ svc *logicpriority.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("priority_create",
		mcp.WithDescription("Create a new issue priority (requires Administrator rights)."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Priority definition, e.g. {\"key\": \"one\", \"name\": {\"ru\": \"...\", \"en\": \"...\"}, \"order\": 60, \"description\": \"...\"}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
