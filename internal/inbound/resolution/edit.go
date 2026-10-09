package resolution

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicresolution "github.com/hoggerty/yandex-tracker-mcp/internal/logic/resolution"
)

type editTool struct{ svc *logicresolution.Service }

func (t *editTool) Description() mcp.Tool {
	return mcp.NewTool("resolution_edit",
		mcp.WithDescription("Edit an issue resolution (requires Administrator rights)."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Resolution id or key")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Resolution's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": {\"ru\": \"...\", \"en\": \"...\"}, \"order\": 350}")),
	)
}

func (t *editTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
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
	return result.JSON(t.svc.Edit(ctx, id, version, fields))
}
