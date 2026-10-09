package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type checklistItemDeleteTool struct{ svc *logicissue.Service }

func (t *checklistItemDeleteTool) Description() mcp.Tool {
	return mcp.NewTool("issue_checklist_item_delete",
		mcp.WithDescription("Delete a single item from a Yandex Tracker issue's checklist."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("item_id", mcp.Required(), mcp.Description("Checklist item id")),
	)
}

func (t *checklistItemDeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	itemID, err := req.RequireString("item_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.ChecklistItemDelete(ctx, issueKey, itemID))
}
