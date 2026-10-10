package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type countTool struct{ svc *logicissue.Service }

func (t *countTool) Description() mcp.Tool {
	return mcp.NewTool("issue_count",
		mcp.WithDescription("Count Yandex Tracker issues matching a query or filter."),
		mcp.WithString("query", mcp.Description("Tracker query language expression")),
		mcp.WithObject("filter", mcp.Description("Structured filter object")),
	)
}

func (t *countTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	filter, _ := args["filter"].(map[string]any)
	return result.JSON(t.svc.Count(ctx, req.GetString("query", ""), filter))
}
