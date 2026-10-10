package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/component"
)

type GetTool struct{ Svc *logiccomponent.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("component_get",
		mcp.WithDescription("Get the parameters of a single Yandex Tracker component."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Component id")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, id))
}
