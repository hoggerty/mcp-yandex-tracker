package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/issue"
)

type CommentEditTool struct{ Svc *logicissue.Service }

func (t *CommentEditTool) Description() mcp.Tool {
	return mcp.NewTool("issue_comment_edit",
		mcp.WithDescription("Change the text of an existing Yandex Tracker issue comment."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("comment_id", mcp.Required(), mcp.Description("Comment id")),
		mcp.WithString("text", mcp.Required(), mcp.Description("New comment text")),
		mcp.WithArray("summonees", mcp.Description("Logins or ids of users to notify"), mcp.WithStringItems()),
	)
}

func (t *CommentEditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	commentID, err := req.RequireString("comment_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	text, err := req.RequireString("text")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.CommentEdit(ctx, issueKey, commentID, text, req.GetStringSlice("summonees", nil)))
}
