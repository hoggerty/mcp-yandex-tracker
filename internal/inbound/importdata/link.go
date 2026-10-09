package importdata

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicimportdata "github.com/hoggerty/yandex-tracker-mcp/internal/logic/importdata"
)

type linkTool struct{ svc *logicimportdata.Service }

func (t *linkTool) Description() mcp.Tool {
	return mcp.NewTool("import_link",
		mcp.WithDescription("Add a link to an issue preserving historical metadata. Administrator only."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("issue_id", mcp.Required(), mcp.Description("Issue id or key")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"relationship\": \"relates\", \"issue\": \"TEST-2\", \"createdAt\": \"...\", \"createdBy\": 11234567, \"updatedAt\": \"...\", \"updatedBy\": 11234567}; relationship, issue, createdAt, createdBy are required")),
	)
}

func (t *linkTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueID, err := req.RequireString("issue_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Link(ctx, issueID, fields))
}
