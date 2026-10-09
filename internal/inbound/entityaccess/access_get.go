package entityaccess

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentityaccess "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityaccess"
)

type accessGetTool struct{ svc *logicentityaccess.Service }

func (t *accessGetTool) Description() mcp.Tool {
	return mcp.NewTool("entityaccess_access_get",
		mcp.WithDescription("Get the extended permissions (ACL grants and permission sources) for a Yandex Tracker entity."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
	)
}

func (t *accessGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.AccessGet(ctx, entityType, id))
}
