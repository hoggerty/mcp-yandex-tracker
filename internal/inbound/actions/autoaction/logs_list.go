package autoaction

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicautoaction "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/autoaction"
)

type LogsListTool struct{ Svc *logicautoaction.Service }

func (t *LogsListTool) Description() mcp.Tool {
	return mcp.NewTool("autoaction_logs_list",
		mcp.WithDescription("List the run history of a Yandex Tracker auto action."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("autoaction_id", mcp.Required(), mcp.Description("Auto action id")),
	)
}

func (t *LogsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	autoActionID, err := req.RequireString("autoaction_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.LogsList(ctx, queueKey, autoActionID))
}
