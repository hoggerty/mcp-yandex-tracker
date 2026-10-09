package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type attachmentListTool struct{ svc *logicissue.Service }

func (t *attachmentListTool) Description() mcp.Tool {
	return mcp.NewTool("issue_attachment_list",
		mcp.WithDescription("List the files attached to a Yandex Tracker issue."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
	)
}

func (t *attachmentListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.AttachmentList(ctx, issueKey))
}
