package bulkchange

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicbulkchange "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/bulkchange"
)

type UpdateTool struct{ Svc *logicbulkchange.Service }

func (t *UpdateTool) Description() mcp.Tool {
	return mcp.NewTool("bulkchange_update",
		mcp.WithDescription("Start a background operation that sets field values on many Yandex Tracker issues (up to 10000 issues)."),
		mcp.WithArray("issues", mcp.Required(), mcp.Description("Issue keys to update"), mcp.WithStringItems()),
		mcp.WithObject("values", mcp.Required(), mcp.Description("Field values to set, e.g. {\"priority\": \"minor\"}")),
		mcp.WithBoolean("notify", mcp.Description("Notify affected users (default false)")),
	)
}

func (t *UpdateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issues, err := req.RequireStringSlice("issues")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	values, ok := req.GetArguments()["values"].(map[string]any)
	if !ok || len(values) == 0 {
		return mcp.NewToolResultError("values must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Update(ctx, issues, values, req.GetBool("notify", false)))
}
