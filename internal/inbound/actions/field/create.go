package field

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicfield "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/field"
)

type CreateTool struct{ Svc *logicfield.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("field_create",
		mcp.WithDescription("Create a new global issue field."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Field definition, e.g. {\"id\": \"myField\", \"name\": {\"en\": \"My Field\", \"ru\": \"Моё поле\"}, \"category\": \"<category id>\", \"type\": \"ru.yandex.startrek.core.fields.StringFieldType\"}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, fields))
}
