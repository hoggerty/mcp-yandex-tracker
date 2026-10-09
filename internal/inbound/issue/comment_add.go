package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type commentAddTool struct{ svc *logicissue.Service }

func (t *commentAddTool) Description() mcp.Tool {
	return mcp.NewTool("issue_comment_add",
		mcp.WithDescription("Add a comment to a Yandex Tracker issue, optionally summoning other users."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("text", mcp.Required(), mcp.Description("Comment body")),
		mcp.WithArray("summonees", mcp.Description("Logins of users to notify/summon")),
	)
}

func (t *commentAddTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	text, err := req.RequireString("text")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	summonees := req.GetStringSlice("summonees", nil)
	return result.JSON(t.svc.AddComment(ctx, issueKey, text, summonees))
}
