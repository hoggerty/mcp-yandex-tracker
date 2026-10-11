package field

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicfield "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/field"
)

type CategoryEditTool struct{ Svc *logicfield.Service }

func (t *CategoryEditTool) Description() mcp.Tool {
	return mcp.NewTool("field_category_edit",
		mcp.WithDescription("Edit a global-field category's name, description, or order."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Category id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Category's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": {\"en\": \"...\", \"ru\": \"...\"}, \"order\": 400}")),
	)
}

func (t *CategoryEditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	return result.JSON(t.Svc.CategoryEdit(ctx, id, version, fields))
}
