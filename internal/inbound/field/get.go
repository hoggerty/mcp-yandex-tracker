package field

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicfield "github.com/hoggerty/yandex-tracker-mcp/internal/logic/field"
)

type getTool struct{ svc *logicfield.Service }

func (t *getTool) Description() mcp.Tool {
	return mcp.NewTool("field_get",
		mcp.WithDescription("Get the parameters of a single global issue field."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("id", mcp.Required(), mcp.Description("Field id/key")),
	)
}

func (t *getTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Get(ctx, id))
}
