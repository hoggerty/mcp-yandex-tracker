package field

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicfield "github.com/hoggerty/yandex-tracker-mcp/internal/logic/field"
)

type listTool struct{ svc *logicfield.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("field_list",
		mcp.WithDescription("List all global issue fields of the organization."),
		mcp.WithReadOnlyHintAnnotation(true),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx))
}
