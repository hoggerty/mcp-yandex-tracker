package entityaccess

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicentityaccess "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/entityaccess"
)

type AccessGetTool struct{ Svc *logicentityaccess.Service }

func (t *AccessGetTool) Description() mcp.Tool {
	return mcp.NewTool("entityaccess_access_get",
		mcp.WithDescription("Get the extended permissions (ACL grants and permission sources) for a Yandex Tracker entity."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
	)
}

func (t *AccessGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.AccessGet(ctx, entityType, id))
}
