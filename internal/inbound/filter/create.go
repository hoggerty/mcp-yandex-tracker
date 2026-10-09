package filter

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicfilter "github.com/hoggerty/yandex-tracker-mcp/internal/logic/filter"
)

type createTool struct{ svc *logicfilter.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("filter_create",
		mcp.WithDescription("Create a new saved issue filter."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Filter definition, e.g. {\"name\": \"My Open Issues\", \"filter\": {\"assignee\": \"me()\", \"status\": \"open\"}}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
