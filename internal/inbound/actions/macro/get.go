package macro

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicmacro "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/macro"
)

type GetTool struct{ Svc *logicmacro.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("macro_get",
		mcp.WithDescription("Get the parameters of a single Yandex Tracker macro."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("macro_id", mcp.Required(), mcp.Description("Macro id")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	macroID, err := req.RequireString("macro_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, queueKey, macroID))
}
