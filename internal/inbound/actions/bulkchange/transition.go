package bulkchange

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicbulkchange "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/bulkchange"
)

type TransitionTool struct{ Svc *logicbulkchange.Service }

func (t *TransitionTool) Description() mcp.Tool {
	return mcp.NewTool("bulkchange_transition",
		mcp.WithDescription("Start a background operation that executes a workflow transition on many Yandex Tracker issues (up to 10000 issues)."),
		mcp.WithString("transition", mcp.Required(), mcp.Description("Transition id, see issue_transitions_list")),
		mcp.WithArray("issues", mcp.Required(), mcp.Description("Issue keys to transition"), mcp.WithStringItems()),
		mcp.WithObject("values", mcp.Description("Additional field values to set after the transition")),
		mcp.WithBoolean("notify", mcp.Description("Notify affected users (default false)")),
	)
}

func (t *TransitionTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	transition, err := req.RequireString("transition")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	issues, err := req.RequireStringSlice("issues")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	values, _ := req.GetArguments()["values"].(map[string]any)
	return result.JSON(t.Svc.Transition(ctx, transition, issues, values, req.GetBool("notify", false)))
}
