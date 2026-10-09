package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type linkDeleteTool struct{ svc *logicissue.Service }

func (t *linkDeleteTool) Description() mcp.Tool {
	return mcp.NewTool("issue_link_delete",
		mcp.WithDescription("Remove a link between two Yandex Tracker issues."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("link_id", mcp.Required(), mcp.Description("Link id")),
	)
}

func (t *linkDeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	linkID, err := req.RequireString("link_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.LinkDelete(ctx, issueKey, linkID))
}
