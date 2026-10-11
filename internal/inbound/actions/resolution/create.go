package resolution

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicresolution "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/resolution"
)

type CreateTool struct{ Svc *logicresolution.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("resolution_create",
		mcp.WithDescription("Create a new issue resolution (requires Administrator rights)."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Resolution definition, e.g. {\"key\": \"myResolution\", \"name\": {\"ru\": \"...\", \"en\": \"...\"}}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, fields))
}
