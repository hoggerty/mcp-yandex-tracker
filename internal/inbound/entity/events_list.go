package entity

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entity"
)

type eventsListTool struct{ svc *logicentity.Service }

func (t *eventsListTool) Description() mcp.Tool {
	return mcp.NewTool("entity_events_list",
		mcp.WithDescription("Get a Yandex Tracker entity's change history, paginated relative to an event id."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithNumber("per_page", mcp.Description("Max events per response, default 50")),
		mcp.WithString("from", mcp.Description("Event id to start after")),
		mcp.WithString("selected", mcp.Description("Event id to center results around")),
	)
}

func (t *eventsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.EventsList(ctx, entityType, id, req.GetInt("per_page", 0), req.GetString("from", ""), req.GetString("selected", "")))
}
