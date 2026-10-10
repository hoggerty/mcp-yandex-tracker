package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type searchTool struct{ svc *logicissue.Service }

func (t *searchTool) Description() mcp.Tool {
	return mcp.NewTool("issue_search",
		mcp.WithDescription("Search Yandex Tracker issues using either a query-language string or a structured filter. At least one of query/filter must be set."),
		mcp.WithString("query", mcp.Description("Tracker query language expression, e.g. \"Queue: QUEUE AND Status: Open\"")),
		mcp.WithObject("filter", mcp.Description("Structured filter, e.g. {\"queue\": \"QUEUE\", \"status\": \"open\"}")),
		mcp.WithString("order_by", mcp.Description("Sort order, e.g. \"+updated\" or \"-created\"")),
		mcp.WithNumber("per_page", mcp.Description("Results per page (default 50)")),
		mcp.WithNumber("page", mcp.Description("Page number, starting at 1")),
	)
}

func (t *searchTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	filter, _ := args["filter"].(map[string]any)
	return result.JSON(t.svc.Search(ctx, req.GetString("query", ""),
		filter,
		req.GetString("order_by", ""),
		req.GetInt("per_page", 0),
		req.GetInt("page", 0)))
}
