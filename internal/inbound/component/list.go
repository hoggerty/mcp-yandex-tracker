package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
)

type listTool struct{ svc *logiccomponent.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("component_list",
		mcp.WithDescription("List all components in the organization."),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx))
}
