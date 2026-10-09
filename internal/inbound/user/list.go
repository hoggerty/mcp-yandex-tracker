package user

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicuser "github.com/hoggerty/yandex-tracker-mcp/internal/logic/user"
)

type listTool struct{ svc *logicuser.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("users_list",
		mcp.WithDescription("List users known to the Yandex Tracker organization."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithNumber("per_page", mcp.Description("Results per page")),
		mcp.WithNumber("page", mcp.Description("Page number, starting at 1")),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx, req.GetInt("per_page", 0), req.GetInt("page", 0)))
}
