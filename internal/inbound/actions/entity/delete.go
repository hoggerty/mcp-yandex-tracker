package entity

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicentity "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/entity"
)

type DeleteTool struct{ Svc *logicentity.Service }

func (t *DeleteTool) Description() mcp.Tool {
	return mcp.NewTool("entity_delete",
		mcp.WithDescription("Delete a Yandex Tracker entity (project, portfolio, or goal)."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithBoolean("with_board", mcp.Description("Also delete the entity's board; default false")),
	)
}

func (t *DeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Delete(ctx, entityType, id, req.GetBool("with_board", false)))
}
