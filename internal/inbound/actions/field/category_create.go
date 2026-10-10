package field

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicfield "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/field"
)

type CategoryCreateTool struct{ Svc *logicfield.Service }

func (t *CategoryCreateTool) Description() mcp.Tool {
	return mcp.NewTool("field_category_create",
		mcp.WithDescription("Create a new category for grouping global issue fields."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Category definition, e.g. {\"name\": {\"en\": \"...\", \"ru\": \"...\"}, \"order\": 400}")),
	)
}

func (t *CategoryCreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.CategoryCreate(ctx, fields))
}
