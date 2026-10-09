package autoaction

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicautoaction "github.com/hoggerty/yandex-tracker-mcp/internal/logic/autoaction"
)

type logGetTool struct{ svc *logicautoaction.Service }

func (t *logGetTool) Description() mcp.Tool {
	return mcp.NewTool("autoaction_log_get",
		mcp.WithDescription("Get a single auto action run, including the issues it touched."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("autoaction_id", mcp.Required(), mcp.Description("Auto action id")),
		mcp.WithString("run_id", mcp.Required(), mcp.Description("Auto action run id")),
	)
}

func (t *logGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	autoActionID, err := req.RequireString("autoaction_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	runID, err := req.RequireString("run_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.LogGet(ctx, queueKey, autoActionID, runID))
}
