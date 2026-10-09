package resolution

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicresolution "github.com/hoggerty/yandex-tracker-mcp/internal/logic/resolution"
)

type createTool struct{ svc *logicresolution.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("resolution_create",
		mcp.WithDescription("Create a new issue resolution (requires Administrator rights)."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Resolution definition, e.g. {\"key\": \"myResolution\", \"name\": {\"ru\": \"...\", \"en\": \"...\"}}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
