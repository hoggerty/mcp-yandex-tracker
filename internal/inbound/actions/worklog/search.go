package worklog

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicworklog "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/worklog"
)

type SearchTool struct{ Svc *logicworklog.Service }

func (t *SearchTool) Description() mcp.Tool {
	return mcp.NewTool("worklog_search",
		mcp.WithDescription("Search worklog records across all issues by author and/or creation date range."),
		mcp.WithString("created_by", mcp.Description("Filter by author login or id")),
		mcp.WithString("created_from", mcp.Description("Filter by creation date, from (ISO 8601)")),
		mcp.WithString("created_to", mcp.Description("Filter by creation date, to (ISO 8601)")),
	)
}

func (t *SearchTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.Search(ctx, req.GetString("created_by", ""),
		req.GetString("created_from", ""),
		req.GetString("created_to", "")))
}
