package worklog

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicworklog "github.com/hoggerty/yandex-tracker-mcp/internal/logic/worklog"
)

type addTool struct{ svc *logicworklog.Service }

func (t *addTool) Description() mcp.Tool {
	return mcp.NewTool("worklog_add",
		mcp.WithDescription("Record time spent on a Yandex Tracker issue."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("start", mcp.Required(), mcp.Description("Work start timestamp, ISO 8601, e.g. 2021-09-21T15:30:00.000+0000")),
		mcp.WithString("duration", mcp.Required(), mcp.Description("Time spent, ISO 8601 duration, e.g. P1DT4H")),
		mcp.WithString("comment", mcp.Description("Optional comment")),
	)
}

func (t *addTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	start, err := req.RequireString("start")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	duration, err := req.RequireString("duration")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Add(ctx, issueKey, start, duration, req.GetString("comment", "")))
}
