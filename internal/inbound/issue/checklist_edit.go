package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type checklistEditTool struct{ svc *logicissue.Service }

func (t *checklistEditTool) Description() mcp.Tool {
	return mcp.NewTool("issue_checklist_edit",
		mcp.WithDescription("Bulk-edit items of a Yandex Tracker issue's checklist; each item must include its id."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithArray("items", mcp.Required(), mcp.Description("Checklist items to update, each an object with at least an \"id\" field")),
	)
}

func (t *checklistEditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	raw, _ := req.GetArguments()["items"].([]any)
	if len(raw) == 0 {
		return mcp.NewToolResultError("items must be a non-empty array"), nil
	}
	items := make([]map[string]any, 0, len(raw))
	for _, v := range raw {
		item, ok := v.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("items must be an array of objects"), nil
		}
		items = append(items, item)
	}
	return result.JSON(t.svc.ChecklistEdit(ctx, issueKey, items))
}
