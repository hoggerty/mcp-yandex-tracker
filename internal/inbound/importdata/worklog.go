package importdata

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicimportdata "github.com/hoggerty/yandex-tracker-mcp/internal/logic/importdata"
)

type worklogTool struct{ svc *logicimportdata.Service }

func (t *worklogTool) Description() mcp.Tool {
	return mcp.NewTool("import_worklog",
		mcp.WithDescription("Add a worklog entry to an issue preserving historical metadata. Administrator only."),
		mcp.WithString("issue_id", mcp.Required(), mcp.Description("Issue id or key")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"duration\": \"PT1H\", \"createdAt\": \"...\", \"createdBy\": \"login\", \"start\": \"...\", \"comment\": \"...\"}")),
	)
}

func (t *worklogTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueID, err := req.RequireString("issue_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Worklog(ctx, issueID, fields))
}
