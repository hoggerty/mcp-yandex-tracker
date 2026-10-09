package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type commentReactionAddTool struct{ svc *logicissue.Service }

func (t *commentReactionAddTool) Description() mcp.Tool {
	return mcp.NewTool("issue_comment_reaction_add",
		mcp.WithDescription("Add a reaction to a Yandex Tracker issue comment."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("comment_id", mcp.Required(), mcp.Description("Comment id")),
		mcp.WithString("reaction", mcp.Required(), mcp.Description("One of LIKE, DISLIKE, LAUGH, HOORAY, CONFUSED, HEART, ROCKET, EYES, FIRE, OK, FACEPALM, CHECK")),
	)
}

func (t *commentReactionAddTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	commentID, err := req.RequireString("comment_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	reaction, err := req.RequireString("reaction")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.CommentReactionAdd(ctx, issueKey, commentID, reaction))
}
