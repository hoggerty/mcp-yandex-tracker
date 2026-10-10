package gap

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicgap "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/gap"
)

type CreateTool struct{ Svc *logicgap.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("gap_create",
		mcp.WithDescription("Create employee absence records (vacation, illness, trip, etc). Administrator only."),
		mcp.WithArray("gaps", mcp.Required(), mcp.Description("Array of absence objects, e.g. [{\"user\": \"login\", \"workflow\": \"vacation\", \"from\": \"2026-07-01T00:00:00.000Z\", \"to\": \"2026-07-15T00:00:00.000Z\", \"fullDay\": true, \"workInAbsence\": false}]; user, workflow, from, to are required")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	raw, ok := req.GetArguments()["gaps"].([]any)
	if !ok || len(raw) == 0 {
		return mcp.NewToolResultError("gaps must be a non-empty array"), nil
	}
	gaps := make([]map[string]any, 0, len(raw))
	for _, v := range raw {
		g, ok := v.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("each gap must be an object"), nil
		}
		gaps = append(gaps, g)
	}
	return result.JSON(t.Svc.Create(ctx, gaps))
}
