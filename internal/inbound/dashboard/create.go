package dashboard

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicdashboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/dashboard"
)

type createTool struct{ svc *logicdashboard.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("dashboard_create",
		mcp.WithDescription("Create a new Yandex Tracker dashboard."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Dashboard definition, e.g. {\"name\": \"...\", \"description\": \"...\"}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
