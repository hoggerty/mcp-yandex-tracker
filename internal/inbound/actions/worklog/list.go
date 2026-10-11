package worklog

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicworklog "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/worklog"
)

type ListTool struct{ Svc *logicworklog.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("worklog_list",
		mcp.WithDescription("List all worklog records for a Yandex Tracker issue."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.List(ctx, issueKey))
}
