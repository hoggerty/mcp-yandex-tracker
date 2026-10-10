package component

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logiccomponent "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/component"
)

type CreateTool struct{ Svc *logiccomponent.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("component_create",
		mcp.WithDescription("Create a new component on a Yandex Tracker queue."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Component definition, e.g. {\"name\": \"...\", \"queue\": \"TEST\", \"description\": \"...\", \"lead\": \"login\", \"assignAuto\": false}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, fields))
}
