package filter

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicfilter "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/filter"
)

type UpdateTool struct{ Svc *logicfilter.Service }

func (t *UpdateTool) Description() mcp.Tool {
	return mcp.NewTool("filter_update",
		mcp.WithDescription("Replace a saved issue filter's name, conditions, and/or sort order. The filter conditions are replaced entirely, not merged — include all desired conditions."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Filter id")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Fields to set, e.g. {\"name\": \"...\", \"filter\": {\"status\": \"open\"}, \"sorts\": [{\"field\": \"priority\", \"isAscending\": false}]}")),
	)
}

func (t *UpdateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Update(ctx, id, fields))
}
