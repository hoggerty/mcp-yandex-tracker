package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/issue"
)

type TransitionExecuteTool struct{ Svc *logicissue.Service }

func (t *TransitionExecuteTool) Description() mcp.Tool {
	return mcp.NewTool("issue_transition_execute",
		mcp.WithDescription("Execute a workflow transition on an issue, moving it to a new status."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("transition_id", mcp.Required(), mcp.Description("Transition ID, as returned by issue_transitions_list")),
		mcp.WithObject("fields", mcp.Description("Optional fields required by the transition's screen (e.g. resolution)")),
	)
}

func (t *TransitionExecuteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	transitionID, err := req.RequireString("transition_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, _ := req.GetArguments()["fields"].(map[string]any)
	return result.JSON(t.Svc.ExecuteTransition(ctx, issueKey, transitionID, fields))
}
