package project

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicproject "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/project"
)

type ListTool struct{ Svc *logicproject.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("project_list",
		mcp.WithDescription("List all Yandex Tracker projects."),
		mcp.WithString("expand", mcp.Description("Optional: \"queues\"")),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx, req.GetString("expand", "")))
}
