package filter

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicfilter "github.com/hoggerty/yandex-tracker-mcp/internal/logic/filter"
)

type deleteTool struct{ svc *logicfilter.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("filter_delete",
		mcp.WithDescription("Delete a saved issue filter."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("id", mcp.Required(), mcp.Description("Filter id")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, id))
}
