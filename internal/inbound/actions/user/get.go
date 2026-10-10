package user

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicuser "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/user"
)

type GetTool struct{ Svc *logicuser.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("user_get",
		mcp.WithDescription("Get a single Yandex Tracker user by login or numeric id."),
		mcp.WithString("id", mcp.Required(), mcp.Description("User login or numeric id (for a purely numeric login, use \"login:12345\")")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, id))
}
