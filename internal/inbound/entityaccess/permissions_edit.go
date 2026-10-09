package entityaccess

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentityaccess "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityaccess"
)

type permissionsEditTool struct{ svc *logicentityaccess.Service }

func (t *permissionsEditTool) Description() mcp.Tool {
	return mcp.NewTool("entityaccess_permissions_edit",
		mcp.WithDescription("Grant or revoke plain permissions (users/groups/roles) on a Yandex Tracker entity."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithObject("body", mcp.Required(), mcp.Description("e.g. {\"grant\": [{\"type\": \"user\", \"id\": \"...\"}], \"revoke\": [...]}")),
	)
}

func (t *permissionsEditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	body, ok := req.GetArguments()["body"].(map[string]any)
	if !ok || len(body) == 0 {
		return mcp.NewToolResultError("body must be a non-empty object"), nil
	}
	return result.JSON(t.svc.PermissionsEdit(ctx, entityType, id, body))
}
