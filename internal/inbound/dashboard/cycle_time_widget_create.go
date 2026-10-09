package dashboard

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicdashboard "github.com/hoggerty/yandex-tracker-mcp/internal/logic/dashboard"
)

type cycleTimeWidgetCreateTool struct{ svc *logicdashboard.Service }

func (t *cycleTimeWidgetCreateTool) Description() mcp.Tool {
	return mcp.NewTool("dashboard_cycle_time_widget_create",
		mcp.WithDescription("Create a \"Cycle time\" widget on a Yandex Tracker dashboard."),
		mcp.WithString("dashboard_id", mcp.Required(), mcp.Description("Dashboard id")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Widget definition, e.g. {\"name\": \"...\", \"filter\": {...}}")),
	)
}

func (t *cycleTimeWidgetCreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	dashboardID, err := req.RequireString("dashboard_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.CreateCycleTimeWidget(ctx, dashboardID, fields))
}
