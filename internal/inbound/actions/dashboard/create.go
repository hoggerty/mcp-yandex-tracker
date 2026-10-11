package dashboard

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicdashboard "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/dashboard"
)

type CreateTool struct{ Svc *logicdashboard.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("dashboard_create",
		mcp.WithDescription("Create a new Yandex Tracker dashboard."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Dashboard definition, e.g. {\"name\": \"...\", \"description\": \"...\"}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, fields))
}
