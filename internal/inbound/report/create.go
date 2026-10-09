package report

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicreport "github.com/hoggerty/yandex-tracker-mcp/internal/logic/report"
)

type createTool struct{ svc *logicreport.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("report_create",
		mcp.WithDescription("Generate a Yandex Tracker issue report (e.g. an xlsx/csv/xml export of a filtered issue list)."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Report fields, e.g. {\"summary\": \"...\", \"parameters\": {\"type\": \"issueFilterExport\", \"format\": \"xlsx\", \"filter\": {\"query\": \"...\"}}}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
