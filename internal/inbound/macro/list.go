package macro

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicmacro "github.com/hoggerty/yandex-tracker-mcp/internal/logic/macro"
)

type listTool struct{ svc *logicmacro.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("macro_list",
		mcp.WithDescription("List the macros configured on a Yandex Tracker queue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.List(ctx, queueKey))
}
