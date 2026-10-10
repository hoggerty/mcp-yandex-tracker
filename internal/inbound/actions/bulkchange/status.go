package bulkchange

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicbulkchange "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/bulkchange"
)

type StatusTool struct{ Svc *logicbulkchange.Service }

func (t *StatusTool) Description() mcp.Tool {
	return mcp.NewTool("bulkchange_status",
		mcp.WithDescription("Get the status of a background bulk-change operation."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Bulk-change operation id")),
	)
}

func (t *StatusTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Status(ctx, id))
}
