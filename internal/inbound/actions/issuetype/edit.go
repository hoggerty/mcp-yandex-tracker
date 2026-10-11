package issuetype

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicissuetype "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/issuetype"
)

type EditTool struct{ Svc *logicissuetype.Service }

func (t *EditTool) Description() mcp.Tool {
	return mcp.NewTool("issuetype_edit",
		mcp.WithDescription("Edit an issue type (requires Administrator rights)."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Issue type id or key")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Issue type's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": {\"ru\": \"...\", \"en\": \"...\"}}")),
	)
}

func (t *EditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	return result.JSON(t.Svc.Edit(ctx, id, version, fields))
}
