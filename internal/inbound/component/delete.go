package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
)

type deleteTool struct{ svc *logiccomponent.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("component_delete",
		mcp.WithDescription("Delete a Yandex Tracker component."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Component id")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, id))
}
