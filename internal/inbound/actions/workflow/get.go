package workflow

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicworkflow "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/workflow"
)

type GetTool struct{ Svc *logicworkflow.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("workflow_get",
		mcp.WithDescription("Get a single Yandex Tracker workflow by id."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Workflow id")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, id))
}
