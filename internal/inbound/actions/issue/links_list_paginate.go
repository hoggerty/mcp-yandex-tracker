package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicissue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/issue"
)

type LinksListPaginateTool struct{ Svc *logicissue.Service }

func (t *LinksListPaginateTool) Description() mcp.Tool {
	return mcp.NewTool("issue_links_list_paginate",
		mcp.WithDescription("List a Yandex Tracker issue's links a page at a time, optionally filtered by link type."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithNumber("per_page", mcp.Description("Links per page")),
		mcp.WithNumber("page", mcp.Description("Page number, starting at 1")),
		mcp.WithArray("link_types", mcp.Description("Link type ids to filter by, e.g. [\"relates\"]"), mcp.WithStringItems()),
	)
}

func (t *LinksListPaginateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.LinksListPaginate(ctx, issueKey,
		req.GetInt("per_page", 0),
		req.GetInt("page", 0),
		req.GetStringSlice("link_types", nil)))
}
