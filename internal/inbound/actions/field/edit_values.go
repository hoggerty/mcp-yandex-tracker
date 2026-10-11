package field

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicfield "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/field"
)

type EditValuesTool struct{ Svc *logicfield.Service }

func (t *EditValuesTool) Description() mcp.Tool {
	return mcp.NewTool("field_edit_values",
		mcp.WithDescription("Edit a global issue field's parameters: category, order, description, visibility, or possible values."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Field id/key")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Field's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"category\": \"...\", \"order\": 102, \"optionsProvider\": {\"type\": \"FixedListOptionsProvider\", \"values\": [\"a\", \"b\"]}}")),
	)
}

func (t *EditValuesTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	return result.JSON(t.Svc.EditValues(ctx, id, version, fields))
}
