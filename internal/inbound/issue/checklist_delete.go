package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type checklistDeleteTool struct{ svc *logicissue.Service }

func (t *checklistDeleteTool) Description() mcp.Tool {
	return mcp.NewTool("issue_checklist_delete",
		mcp.WithDescription("Delete a Yandex Tracker issue's entire checklist."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
	)
}

func (t *checklistDeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.ChecklistDelete(ctx, issueKey))
}
