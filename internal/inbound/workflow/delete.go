package workflow

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicworkflow "github.com/hoggerty/yandex-tracker-mcp/internal/logic/workflow"
)

type deleteTool struct{ svc *logicworkflow.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("workflow_delete",
		mcp.WithDescription("Delete a Yandex Tracker workflow."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Workflow id")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, id))
}
