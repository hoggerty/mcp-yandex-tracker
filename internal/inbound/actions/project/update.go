package project

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicproject "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/project"
)

type UpdateTool struct{ Svc *logicproject.Service }

func (t *UpdateTool) Description() mcp.Tool {
	return mcp.NewTool("project_update",
		mcp.WithDescription("Replace a Yandex Tracker project's fields."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Project id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Current project version, for optimistic locking")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"name\": \"...\", \"queues\": \"QUEUE_KEY\", \"description\": \"...\", \"lead\": \"login\", \"status\": \"DRAFT|IN_PROGRESS|LAUNCHED|POSTPONED\", \"startDate\": \"YYYY-MM-DD\", \"endDate\": \"YYYY-MM-DD\"}")),
	)
}

func (t *UpdateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	version, err := req.RequireInt("version")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Update(ctx, id, version, fields))
}
