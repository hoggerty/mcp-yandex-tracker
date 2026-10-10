package gap

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicgap "github.com/hoggerty/yandex-tracker-mcp/internal/logic/gap"
)

type deleteTool struct{ svc *logicgap.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("gap_delete",
		mcp.WithDescription("Delete employee absence records by id (up to 100 at once). Administrator only."),
		mcp.WithArray("gap_ids", mcp.Required(), mcp.Description("Absence record ids to delete"), mcp.WithStringItems()),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	raw, ok := req.GetArguments()["gap_ids"].([]any)
	if !ok || len(raw) == 0 {
		return mcp.NewToolResultError("gap_ids must be a non-empty array"), nil
	}
	ids := make([]string, 0, len(raw))
	for _, v := range raw {
		s, ok := v.(string)
		if !ok || s == "" {
			return mcp.NewToolResultError("each gap_id must be a non-empty string"), nil
		}
		ids = append(ids, s)
	}
	return result.JSON(t.svc.Delete(ctx, ids))
}
