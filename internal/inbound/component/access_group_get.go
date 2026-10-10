package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
)

type accessGroupGetTool struct{ svc *logiccomponent.Service }

func (t *accessGroupGetTool) Description() mcp.Tool {
	return mcp.NewTool("component_access_group_get",
		mcp.WithDescription("Get a group's permissions for a Yandex Tracker component."),
		mcp.WithString("component_id", mcp.Required(), mcp.Description("Component id")),
		mcp.WithString("group_id", mcp.Required(), mcp.Description("Group id")),
	)
}

func (t *accessGroupGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	componentID, err := req.RequireString("component_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	groupID, err := req.RequireString("group_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.AccessGroupGet(ctx, componentID, groupID))
}
