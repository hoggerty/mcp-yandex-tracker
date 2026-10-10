package sprint

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicsprint "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/sprint"
)

type ArchiveTool struct{ Svc *logicsprint.Service }

func (t *ArchiveTool) Description() mcp.Tool {
	return mcp.NewTool("sprint_archive",
		mcp.WithDescription("Archive a Yandex Tracker sprint."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Sprint id")),
		mcp.WithNumber("version", mcp.Description("Sprint's current version (for optimistic locking); omit to skip the check")),
	)
}

func (t *ArchiveTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Archive(ctx, id, req.GetInt("version", 0)))
}
