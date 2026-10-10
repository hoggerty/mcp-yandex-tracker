package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/issue"
)

type MoveTool struct{ Svc *logicissue.Service }

func (t *MoveTool) Description() mcp.Tool {
	return mcp.NewTool("issue_move",
		mcp.WithDescription("Move a Yandex Tracker issue to another queue."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Destination queue key")),
		mcp.WithBoolean("notify", mcp.Description("Notify issue participants (default true)")),
		mcp.WithBoolean("notify_author", mcp.Description("Notify the issue author (default false)")),
		mcp.WithBoolean("move_all_fields", mcp.Description("Move field values that don't exist in the destination queue to the description (default false)")),
		mcp.WithBoolean("initial_status", mcp.Description("Reset the issue status to the destination queue's initial status (default false)")),
		mcp.WithObject("fields", mcp.Description("Additional fields to set during the move")),
	)
}

func (t *MoveTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	queue, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, _ := req.GetArguments()["fields"].(map[string]any)
	return result.JSON(t.Svc.Move(ctx, issueKey, queue,
		req.GetBool("notify", true),
		req.GetBool("notify_author", false),
		req.GetBool("move_all_fields", false),
		req.GetBool("initial_status", false),
		fields))
}
