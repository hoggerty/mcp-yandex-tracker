package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicissue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/issue"
)

type CommentGetTool struct{ Svc *logicissue.Service }

func (t *CommentGetTool) Description() mcp.Tool {
	return mcp.NewTool("issue_comment_get",
		mcp.WithDescription("Retrieve a single Yandex Tracker issue comment."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("comment_id", mcp.Required(), mcp.Description("Comment id")),
	)
}

func (t *CommentGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	commentID, err := req.RequireString("comment_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.CommentGet(ctx, issueKey, commentID))
}
