package entity

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entity"
)

type deleteTool struct{ svc *logicentity.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("entity_delete",
		mcp.WithDescription("Delete a Yandex Tracker entity (project, portfolio, or goal)."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithBoolean("with_board", mcp.Description("Also delete the entity's board; default false")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, entityType, id, req.GetBool("with_board", false)))
}
