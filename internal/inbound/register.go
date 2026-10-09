package inbound

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/autoaction"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/board"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/bulkchange"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/column"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/component"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/dashboard"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/entity"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/entityaccess"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/entityattachment"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/entitychecklist"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/entitycomment"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/entitylink"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/externallink"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/field"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/filter"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/gap"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/importdata"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/issue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/issuetype"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/macro"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/priority"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/project"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/queue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/report"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/resolution"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/sprint"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/status"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/trigger"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/user"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/version"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/workflow"
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/worklog"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic"
)

func RegisterTools(s *server.MCPServer, svcs *logic.Services, policy Policy) {
	entries := policy.Apply(allTools(svcs))
	tools := make([]server.ServerTool, len(entries))
	for i, e := range entries {
		tools[i] = e.ServerTool()
	}
	s.AddTools(tools...)
}

func allTools(svcs *logic.Services) []tool.Entry {
	var all []tool.Entry
	all = append(all, issue.Tools(svcs.Issue)...)
	all = append(all, queue.Tools(svcs.Queue)...)
	all = append(all, user.Tools(svcs.User)...)
	all = append(all, board.Tools(svcs.Board)...)
	all = append(all, report.Tools(svcs.Report)...)
	all = append(all, worklog.Tools(svcs.Worklog)...)
	all = append(all, externallink.Tools(svcs.ExternalLink)...)
	all = append(all, bulkchange.Tools(svcs.BulkChange)...)
	all = append(all, field.Tools(svcs.Field)...)
	all = append(all, filter.Tools(svcs.Filter)...)
	all = append(all, workflow.Tools(svcs.Workflow)...)
	all = append(all, autoaction.Tools(svcs.AutoAction)...)
	all = append(all, trigger.Tools(svcs.Trigger)...)
	all = append(all, macro.Tools(svcs.Macro)...)
	all = append(all, component.Tools(svcs.Component)...)
	all = append(all, version.Tools(svcs.Version)...)
	all = append(all, column.Tools(svcs.Column)...)
	all = append(all, sprint.Tools(svcs.Sprint)...)
	all = append(all, dashboard.Tools(svcs.Dashboard)...)
	all = append(all, issuetype.Tools(svcs.IssueType)...)
	all = append(all, status.Tools(svcs.Status)...)
	all = append(all, resolution.Tools(svcs.Resolution)...)
	all = append(all, priority.Tools(svcs.Priority)...)
	all = append(all, entity.Tools(svcs.Entity)...)
	all = append(all, entitycomment.Tools(svcs.EntityComment)...)
	all = append(all, entitychecklist.Tools(svcs.EntityChecklist)...)
	all = append(all, entityattachment.Tools(svcs.EntityAttachment)...)
	all = append(all, entitylink.Tools(svcs.EntityLink)...)
	all = append(all, entityaccess.Tools(svcs.EntityAccess)...)
	all = append(all, project.Tools(svcs.Project)...)
	all = append(all, importdata.Tools(svcs.ImportData)...)
	all = append(all, gap.Tools(svcs.Gap)...)
	return all
}
