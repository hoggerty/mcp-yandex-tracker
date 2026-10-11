package report

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicreport "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/report"
)

type SearchTool struct{ Svc *logicreport.Service }

func (t *SearchTool) Description() mcp.Tool {
	return mcp.NewTool("report_search",
		mcp.WithDescription("Find previously generated Yandex Tracker issue reports."),
		mcp.WithObject("filter", mcp.Description("Filter, e.g. {\"author\": \"...\", \"id\": \"...\"}")),
		mcp.WithNumber("per_page", mcp.Description("Reports per page (default 50)")),
		mcp.WithNumber("page", mcp.Description("Page number, starting at 1")),
	)
}

func (t *SearchTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	filter, _ := req.GetArguments()["filter"].(map[string]any)
	return result.JSON(t.Svc.Search(ctx, filter, req.GetInt("per_page", 0), req.GetInt("page", 0)))
}
