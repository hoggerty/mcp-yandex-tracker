package workflow

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicworkflow "github.com/hoggerty/yandex-tracker-mcp/internal/logic/workflow"
)

type listTool struct{ svc *logicworkflow.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("workflow_list",
		mcp.WithDescription("List all Yandex Tracker workflows in the organization."),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx))
}
