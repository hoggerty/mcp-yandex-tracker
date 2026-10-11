package workflow

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicworkflow "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/workflow"
)

type ListTool struct{ Svc *logicworkflow.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("workflow_list",
		mcp.WithDescription("List all Yandex Tracker workflows in the organization."),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx))
}
