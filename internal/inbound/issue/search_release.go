package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type searchReleaseTool struct{ svc *logicissue.Service }

func (t *searchReleaseTool) Description() mcp.Tool {
	return mcp.NewTool("issue_search_release",
		mcp.WithDescription("Release server-side resources held by a scrolled issue_search request."),
		mcp.WithString("scroll_id", mcp.Required(), mcp.Description("Value of the X-Scroll-Id response header from the search")),
		mcp.WithString("scroll_token", mcp.Required(), mcp.Description("Value of the X-Scroll-Token response header from the search")),
	)
}

func (t *searchReleaseTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	scrollID, err := req.RequireString("scroll_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	scrollToken, err := req.RequireString("scroll_token")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.SearchRelease(ctx, scrollID, scrollToken))
}
