package project

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicproject "github.com/hoggerty/yandex-tracker-mcp/internal/logic/project"
)

type listTool struct{ svc *logicproject.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("project_list",
		mcp.WithDescription("List all Yandex Tracker projects."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("expand", mcp.Description("Optional: \"queues\"")),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx, req.GetString("expand", "")))
}
