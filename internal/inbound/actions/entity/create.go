package entity

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicentity "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/entity"
)

type CreateTool struct{ Svc *logicentity.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("entity_create",
		mcp.WithDescription("Create a new Yandex Tracker entity (project, portfolio, or goal)."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Entity definition, e.g. {\"fields\": {\"summary\": \"...\", \"teamAccess\": true}, \"links\": [{\"relationship\": \"works towards\", \"entity\": \"ENTITY_ID\"}]}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, entityType, fields))
}
