package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
)

type accessUserGetTool struct{ svc *logiccomponent.Service }

func (t *accessUserGetTool) Description() mcp.Tool {
	return mcp.NewTool("component_access_user_get",
		mcp.WithDescription("Get a user's permissions for a Yandex Tracker component."),
		mcp.WithString("component_id", mcp.Required(), mcp.Description("Component id")),
		mcp.WithString("user_id", mcp.Required(), mcp.Description("User id or login")),
	)
}

func (t *accessUserGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	componentID, err := req.RequireString("component_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	userID, err := req.RequireString("user_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.AccessUserGet(ctx, componentID, userID))
}
