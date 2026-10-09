package worklog

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicworklog "github.com/hoggerty/yandex-tracker-mcp/internal/logic/worklog"
)

type listTool struct{ svc *logicworklog.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("worklog_list",
		mcp.WithDescription("List all worklog records for a Yandex Tracker issue."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.List(ctx, issueKey))
}
