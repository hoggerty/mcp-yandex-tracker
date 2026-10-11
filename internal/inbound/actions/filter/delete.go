package filter

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicfilter "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/filter"
)

type DeleteTool struct{ Svc *logicfilter.Service }

func (t *DeleteTool) Description() mcp.Tool {
	return mcp.NewTool("filter_delete",
		mcp.WithDescription("Delete a saved issue filter."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Filter id")),
	)
}

func (t *DeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Delete(ctx, id))
}
