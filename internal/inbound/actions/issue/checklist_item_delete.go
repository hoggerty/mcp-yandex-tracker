package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicissue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/issue"
)

type ChecklistItemDeleteTool struct{ Svc *logicissue.Service }

func (t *ChecklistItemDeleteTool) Description() mcp.Tool {
	return mcp.NewTool("issue_checklist_item_delete",
		mcp.WithDescription("Delete a single item from a Yandex Tracker issue's checklist."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("item_id", mcp.Required(), mcp.Description("Checklist item id")),
	)
}

func (t *ChecklistItemDeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	itemID, err := req.RequireString("item_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.ChecklistItemDelete(ctx, issueKey, itemID))
}
