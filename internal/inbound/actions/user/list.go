package user

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicuser "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/user"
)

type ListTool struct{ Svc *logicuser.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("users_list",
		mcp.WithDescription("List users known to the Yandex Tracker organization."),
		mcp.WithNumber("per_page", mcp.Description("Results per page")),
		mcp.WithNumber("page", mcp.Description("Page number, starting at 1")),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx, req.GetInt("per_page", 0), req.GetInt("page", 0)))
}
