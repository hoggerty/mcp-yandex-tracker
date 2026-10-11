package report

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicreport "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/report"
)

type CreateTool struct{ Svc *logicreport.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("report_create",
		mcp.WithDescription("Generate a Yandex Tracker issue report (e.g. an xlsx/csv/xml export of a filtered issue list)."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Report fields, e.g. {\"summary\": \"...\", \"parameters\": {\"type\": \"issueFilterExport\", \"format\": \"xlsx\", \"filter\": {\"query\": \"...\"}}}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, fields))
}
