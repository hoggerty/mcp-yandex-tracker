package workflow

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicworkflow "github.com/hoggerty/yandex-tracker-mcp/internal/logic/workflow"
)

type createTool struct{ svc *logicworkflow.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("workflow_create",
		mcp.WithDescription("Create a new Yandex Tracker workflow."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Workflow definition, e.g. {\"name\": \"Design\", \"queue\": \"DESIGN\", \"initialAction\": {...}, \"steps\": [...]}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
