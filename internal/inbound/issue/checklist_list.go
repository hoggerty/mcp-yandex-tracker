package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type checklistListTool struct{ svc *logicissue.Service }

func (t *checklistListTool) Description() mcp.Tool {
	return mcp.NewTool("issue_checklist_list",
		mcp.WithDescription("List the items of a Yandex Tracker issue's checklist."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
	)
}

func (t *checklistListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.ChecklistList(ctx, issueKey))
}
