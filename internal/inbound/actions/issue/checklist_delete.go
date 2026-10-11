package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicissue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/issue"
)

type ChecklistDeleteTool struct{ Svc *logicissue.Service }

func (t *ChecklistDeleteTool) Description() mcp.Tool {
	return mcp.NewTool("issue_checklist_delete",
		mcp.WithDescription("Delete a Yandex Tracker issue's entire checklist."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
	)
}

func (t *ChecklistDeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.ChecklistDelete(ctx, issueKey))
}
