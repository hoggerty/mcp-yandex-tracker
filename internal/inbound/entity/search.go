package entity

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentity "github.com/hoggerty/yandex-tracker-mcp/internal/logic/entity"
)

type searchTool struct{ svc *logicentity.Service }

func (t *searchTool) Description() mcp.Tool {
	return mcp.NewTool("entity_search",
		mcp.WithDescription("List Yandex Tracker entities of a given type matching a filter."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithObject("body", mcp.Description("e.g. {\"filter\": {\"entityStatus\": \"in_progress\"}, \"orderBy\": \"entityStatus\", \"orderAsc\": true, \"input\": \"substring\"}")),
		mcp.WithString("fields", mcp.Description("Optional comma-separated additional fields to include")),
		mcp.WithNumber("per_page", mcp.Description("Results per page, default 50")),
		mcp.WithNumber("page", mcp.Description("Page number, default 1")),
	)
}

func (t *searchTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	body, _ := req.GetArguments()["body"].(map[string]any)
	return result.JSON(t.svc.Search(ctx, entityType, body, req.GetString("fields", ""), req.GetInt("per_page", 0), req.GetInt("page", 0)))
}
