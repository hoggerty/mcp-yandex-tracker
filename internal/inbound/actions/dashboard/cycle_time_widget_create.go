package dashboard

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicdashboard "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/dashboard"
)

type CycleTimeWidgetCreateTool struct{ Svc *logicdashboard.Service }

func (t *CycleTimeWidgetCreateTool) Description() mcp.Tool {
	return mcp.NewTool("dashboard_cycle_time_widget_create",
		mcp.WithDescription("Create a \"Cycle time\" widget on a Yandex Tracker dashboard."),
		mcp.WithString("dashboard_id", mcp.Required(), mcp.Description("Dashboard id")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Widget definition, e.g. {\"name\": \"...\", \"filter\": {...}}")),
	)
}

func (t *CycleTimeWidgetCreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	dashboardID, err := req.RequireString("dashboard_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.CreateCycleTimeWidget(ctx, dashboardID, fields))
}
