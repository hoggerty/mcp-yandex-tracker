package inbound

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/autoaction"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/board"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/bulkchange"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/column"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/component"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/dashboard"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/entity"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/entityaccess"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/entityattachment"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/entitychecklist"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/entitycomment"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/entitylink"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/externallink"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/field"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/filter"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/gap"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/importdata"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/issue"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/issuetype"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/macro"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/priority"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/project"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/queue"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/report"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/resolution"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/sprint"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/status"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/trigger"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/user"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/version"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/workflow"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/actions/worklog"
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	"github.com/hoggerty/mcp-yandex-tracker/internal/logic"
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
