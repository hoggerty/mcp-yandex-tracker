package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
)

type editTool struct{ svc *logiccomponent.Service }

func (t *editTool) Description() mcp.Tool {
	return mcp.NewTool("component_edit",
		mcp.WithDescription("Edit a Yandex Tracker component."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Component id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Component's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": \"...\", \"description\": \"...\", \"lead\": \"login\", \"assignAuto\": true}")),
	)
}

func (t *editTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	return result.JSON(t.svc.Edit(ctx, id, version, fields))
}
