package logic

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/autoaction"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/board"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/bulkchange"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/column"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/component"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/dashboard"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entity"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entityaccess"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entityattachment"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entitychecklist"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entitycomment"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entitylink"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/externallink"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/field"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/filter"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/gap"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/importdata"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/issue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/issuetype"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/macro"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/priority"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/project"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/queue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/report"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/resolution"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/sprint"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/status"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/trigger"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/user"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/version"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/workflow"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/worklog"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
	"github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"
)

type Services struct {
	Issue            *issue.Service
	Queue            *queue.Service
	User             *user.Service
	Board            *board.Service
	Report           *report.Service
	Worklog          *worklog.Service
	ExternalLink     *externallink.Service
	BulkChange       *bulkchange.Service
	Field            *field.Service
	Filter           *filter.Service
	Workflow         *workflow.Service
	AutoAction       *autoaction.Service
	Trigger          *trigger.Service
	Macro            *macro.Service
	Component        *component.Service
	Version          *version.Service
	Column           *column.Service
	Sprint           *sprint.Service
	Dashboard        *dashboard.Service
	IssueType        *issuetype.Service
	Status           *status.Service
	Resolution       *resolution.Service
	Priority         *priority.Service
	Entity           *entity.Service
	EntityComment    *entitycomment.Service
	EntityChecklist  *entitychecklist.Service
	EntityAttachment *entityattachment.Service
	EntityLink       *entitylink.Service
	EntityAccess     *entityaccess.Service
	Project          *project.Service
	ImportData       *importdata.Service
	Gap              *gap.Service
}

func New(client *trackerapi.Client, guard safety.Guard) *Services {
	return &Services{
		Issue:            issue.New(client, guard),
		Queue:            queue.New(client, guard),
		User:             user.New(client, guard),
		Board:            board.New(client, guard),
		Report:           report.New(client, guard),
		Worklog:          worklog.New(client, guard),
		ExternalLink:     externallink.New(client, guard),
		BulkChange:       bulkchange.New(client, guard),
		Field:            field.New(client, guard),
		Filter:           filter.New(client, guard),
		Workflow:         workflow.New(client, guard),
		AutoAction:       autoaction.New(client, guard),
		Trigger:          trigger.New(client, guard),
		Macro:            macro.New(client, guard),
		Component:        component.New(client, guard),
		Version:          version.New(client, guard),
		Column:           column.New(client, guard),
		Sprint:           sprint.New(client, guard),
		Dashboard:        dashboard.New(client, guard),
		IssueType:        issuetype.New(client, guard),
		Status:           status.New(client, guard),
		Resolution:       resolution.New(client, guard),
		Priority:         priority.New(client, guard),
		Entity:           entity.New(client, guard),
		EntityComment:    entitycomment.New(client, guard),
		EntityChecklist:  entitychecklist.New(client, guard),
		EntityAttachment: entityattachment.New(client, guard),
		EntityLink:       entitylink.New(client, guard),
		EntityAccess:     entityaccess.New(client, guard),
		Project:          project.New(client, guard),
		ImportData:       importdata.New(client, guard),
		Gap:              gap.New(client, guard),
	}
}
