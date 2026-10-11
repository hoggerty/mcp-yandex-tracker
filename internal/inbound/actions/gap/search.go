package gap

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicgap "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/gap"
)

type SearchTool struct{ Svc *logicgap.Service }

func (t *SearchTool) Description() mcp.Tool {
	return mcp.NewTool("gap_search",
		mcp.WithDescription("Find employee absences for given users and date range. Administrator only."),
		mcp.WithObject("body", mcp.Required(), mcp.Description("e.g. {\"users\": [\"login1\", \"login2\"], \"from\": \"2026-07-01T00:00:00.000Z\", \"to\": \"2026-08-31T23:59:59.999Z\"}")),
		mcp.WithNumber("per_page", mcp.Description("Users per page, default 50")),
		mcp.WithNumber("page", mcp.Description("Page number, default 1")),
	)
}

func (t *SearchTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	body, ok := req.GetArguments()["body"].(map[string]any)
	if !ok || len(body) == 0 {
		return mcp.NewToolResultError("body must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Search(ctx, body, req.GetInt("per_page", 0), req.GetInt("page", 0)))
}
