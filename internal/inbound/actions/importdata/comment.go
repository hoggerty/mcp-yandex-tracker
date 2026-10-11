package importdata

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicimportdata "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/importdata"
)

type CommentTool struct{ Svc *logicimportdata.Service }

func (t *CommentTool) Description() mcp.Tool {
	return mcp.NewTool("import_comment",
		mcp.WithDescription("Add a comment to an issue preserving historical metadata. Administrator only."),
		mcp.WithString("issue_id", mcp.Required(), mcp.Description("Issue id or key")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"text\": \"...\", \"createdAt\": \"2017-08-29T12:34:41.740+0000\", \"createdBy\": 11234567, \"updatedAt\": \"...\", \"updatedBy\": 11234567}; text, createdAt, createdBy are required")),
	)
}

func (t *CommentTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueID, err := req.RequireString("issue_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Comment(ctx, issueID, fields))
}
