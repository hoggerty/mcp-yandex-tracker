package autoaction

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicautoaction "github.com/hoggerty/yandex-tracker-mcp/internal/logic/autoaction"
)

type getTool struct{ svc *logicautoaction.Service }

func (t *getTool) Description() mcp.Tool {
	return mcp.NewTool("autoaction_get",
		mcp.WithDescription("Get the parameters of a Yandex Tracker auto action."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("autoaction_id", mcp.Required(), mcp.Description("Auto action id")),
	)
}

func (t *getTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	autoActionID, err := req.RequireString("autoaction_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Get(ctx, queueKey, autoActionID))
}
