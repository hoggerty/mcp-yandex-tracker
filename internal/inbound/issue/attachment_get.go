package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type attachmentGetTool struct{ svc *logicissue.Service }

func (t *attachmentGetTool) Description() mcp.Tool {
	return mcp.NewTool("issue_attachment_get",
		mcp.WithDescription("Retrieve metadata for a single file attached to a Yandex Tracker issue."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("attachment_id", mcp.Required(), mcp.Description("Attachment id")),
	)
}

func (t *attachmentGetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	attachmentID, err := req.RequireString("attachment_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.AttachmentGet(ctx, issueKey, attachmentID))
}
