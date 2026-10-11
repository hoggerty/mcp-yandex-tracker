package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logiccomponent "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/component"
)

type ListTool struct{ Svc *logiccomponent.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("component_list",
		mcp.WithDescription("List all components in the organization."),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx))
}
