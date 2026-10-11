package field

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicfield "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/field"
)

type ListTool struct{ Svc *logicfield.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("field_list",
		mcp.WithDescription("List all global issue fields of the organization."),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx))
}
