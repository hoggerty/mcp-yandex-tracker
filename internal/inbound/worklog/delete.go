package worklog

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicworklog "github.com/hoggerty/yandex-tracker-mcp/internal/logic/worklog"
)

type deleteTool struct{ svc *logicworklog.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("worklog_delete",
		mcp.WithDescription("Delete a worklog record from a Yandex Tracker issue."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("worklog_id", mcp.Required(), mcp.Description("Worklog record id")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	worklogID, err := req.RequireString("worklog_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, issueKey, worklogID))
}
