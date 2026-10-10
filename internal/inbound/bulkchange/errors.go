package bulkchange

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicbulkchange "github.com/hoggerty/yandex-tracker-mcp/internal/logic/bulkchange"
)

type errorsTool struct{ svc *logicbulkchange.Service }

func (t *errorsTool) Description() mcp.Tool {
	return mcp.NewTool("bulkchange_errors",
		mcp.WithDescription("Get the issues that failed within a bulk-change operation, with error details."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Bulk-change operation id")),
	)
}

func (t *errorsTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Errors(ctx, id))
}
