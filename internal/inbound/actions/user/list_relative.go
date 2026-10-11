package user

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicuser "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/user"
)

type ListRelativeTool struct{ Svc *logicuser.Service }

func (t *ListRelativeTool) Description() mcp.Tool {
	return mcp.NewTool("user_list_relative",
		mcp.WithDescription("List Yandex Tracker users, paginated relative to a starting user id (sorted ascending by uid)."),
		mcp.WithNumber("per_page", mcp.Description("Users per page, 1-100, default 50")),
		mcp.WithNumber("id", mcp.Description("User id to start from (use the last uid from the previous page to continue)")),
		mcp.WithString("expand", mcp.Description("Optional: \"groups\" to include group membership")),
	)
}

func (t *ListRelativeTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.ListRelative(ctx, req.GetInt("per_page", 0), req.GetInt("id", 0), req.GetString("expand", "")))
}
