package field

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicfield "github.com/hoggerty/yandex-tracker-mcp/internal/logic/field"
)

type editNameTool struct{ svc *logicfield.Service }

func (t *editNameTool) Description() mcp.Tool {
	return mcp.NewTool("field_edit_name",
		mcp.WithDescription("Rename a global issue field."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Field id/key")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Field's current version (for optimistic locking)")),
		mcp.WithObject("name", mcp.Required(), mcp.Description("New localized name, e.g. {\"en\": \"...\", \"ru\": \"...\"}")),
	)
}

func (t *editNameTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	version, err := req.RequireInt("version")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	name, ok := req.GetArguments()["name"].(map[string]any)
	if !ok || len(name) == 0 {
		return mcp.NewToolResultError("name must be a non-empty object"), nil
	}
	return result.JSON(t.svc.EditName(ctx, id, version, name))
}
