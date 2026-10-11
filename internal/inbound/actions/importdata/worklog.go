package importdata

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicimportdata "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/importdata"
)

type WorklogTool struct{ Svc *logicimportdata.Service }

func (t *WorklogTool) Description() mcp.Tool {
	return mcp.NewTool("import_worklog",
		mcp.WithDescription("Add a worklog entry to an issue preserving historical metadata. Administrator only."),
		mcp.WithString("issue_id", mcp.Required(), mcp.Description("Issue id or key")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"duration\": \"PT1H\", \"createdAt\": \"...\", \"createdBy\": \"login\", \"start\": \"...\", \"comment\": \"...\"}")),
	)
}

func (t *WorklogTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueID, err := req.RequireString("issue_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Worklog(ctx, issueID, fields))
}
