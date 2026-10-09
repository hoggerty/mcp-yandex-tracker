package bulkchange

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicbulkchange "github.com/hoggerty/yandex-tracker-mcp/internal/logic/bulkchange"
)

type moveTool struct{ svc *logicbulkchange.Service }

func (t *moveTool) Description() mcp.Tool {
	return mcp.NewTool("bulkchange_move",
		mcp.WithDescription("Start a background operation that moves many Yandex Tracker issues to another queue (up to 10000 issues)."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Destination queue key")),
		mcp.WithArray("issues", mcp.Required(), mcp.Description("Issue keys to move"), mcp.WithStringItems()),
		mcp.WithObject("values", mcp.Description("Additional field values to set during the move")),
		mcp.WithBoolean("move_all_fields", mcp.Description("Move field values that don't exist in the destination queue to the description (default false)")),
		mcp.WithBoolean("initial_status", mcp.Description("Reset moved issues to the destination queue's initial status (default false)")),
		mcp.WithBoolean("notify", mcp.Description("Notify affected users (default false)")),
	)
}

func (t *moveTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queue, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	issues, err := req.RequireStringSlice("issues")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	values, _ := req.GetArguments()["values"].(map[string]any)
	return result.JSON(t.svc.Move(ctx, queue, issues, values,
		req.GetBool("move_all_fields", false),
		req.GetBool("initial_status", false),
		req.GetBool("notify", false),
	))
}
