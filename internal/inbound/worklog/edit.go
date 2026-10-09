package worklog

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicworklog "github.com/hoggerty/yandex-tracker-mcp/internal/logic/worklog"
)

type editTool struct{ svc *logicworklog.Service }

func (t *editTool) Description() mcp.Tool {
	return mcp.NewTool("worklog_edit",
		mcp.WithDescription("Change the duration and/or comment of an existing worklog record."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("worklog_id", mcp.Required(), mcp.Description("Worklog record id")),
		mcp.WithString("duration", mcp.Required(), mcp.Description("Time spent, ISO 8601 duration, e.g. P1DT4H")),
		mcp.WithString("comment", mcp.Description("Optional comment")),
	)
}

func (t *editTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	worklogID, err := req.RequireString("worklog_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	duration, err := req.RequireString("duration")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Edit(ctx, issueKey, worklogID, duration, req.GetString("comment", "")))
}
