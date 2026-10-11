package project

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicproject "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/project"
)

type CreateTool struct{ Svc *logicproject.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("project_create",
		mcp.WithDescription("Create a new Yandex Tracker project."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"name\": \"...\", \"queues\": \"QUEUE_KEY\", \"description\": \"...\", \"lead\": \"login\", \"status\": \"DRAFT|IN_PROGRESS|LAUNCHED|POSTPONED\", \"startDate\": \"YYYY-MM-DD\", \"endDate\": \"YYYY-MM-DD\"}; name and queues are required")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, fields))
}
