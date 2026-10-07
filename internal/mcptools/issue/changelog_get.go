package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/result"
	"github.com/hoggerty/yandex-tracker-mcp/internal/trackerapi"
)

type changelogGetTool struct{ client *trackerapi.Client }

func (t *changelogGetTool) Description() mcp.Tool {
	return mcp.NewTool("issue_changelog_get",
		mcp.WithDescription("Retrieve the field change history for a Yandex Tracker issue."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
	)
}

func (t *changelogGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.client.GetChangelog(ctx, issueKey))
}
