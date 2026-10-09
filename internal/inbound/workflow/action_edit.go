package workflow

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicworkflow "github.com/hoggerty/yandex-tracker-mcp/internal/logic/workflow"
)

type actionEditTool struct{ svc *logicworkflow.Service }

func (t *actionEditTool) Description() mcp.Tool {
	return mcp.NewTool("workflow_action_edit",
		mcp.WithDescription("Edit a single action of a Yandex Tracker workflow step."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Workflow id")),
		mcp.WithString("status", mcp.Required(), mcp.Description("Status key of the step the action belongs to")),
		mcp.WithString("action_id", mcp.Required(), mcp.Description("Action id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Workflow's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": {\"en\": \"...\"}, \"target\": \"statusKey\", \"conditions\": [], \"functions\": []}")),
	)
}

func (t *actionEditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	status, err := req.RequireString("status")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	actionID, err := req.RequireString("action_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	version, err := req.RequireInt("version")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.ActionEdit(ctx, id, status, actionID, version, fields))
}
