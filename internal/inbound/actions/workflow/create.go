package workflow

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicworkflow "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/workflow"
)

type CreateTool struct{ Svc *logicworkflow.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("workflow_create",
		mcp.WithDescription("Create a new Yandex Tracker workflow."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Workflow definition, e.g. {\"name\": \"Design\", \"queue\": \"DESIGN\", \"initialAction\": {...}, \"steps\": [...]}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, fields))
}
