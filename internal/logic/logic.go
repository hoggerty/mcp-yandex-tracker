package logic

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/autoaction"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/board"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/bulkchange"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/column"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/component"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/dashboard"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/entity"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityaccess"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/entityattachment"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitychecklist"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitycomment"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/entitylink"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/externallink"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/field"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/filter"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/gap"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/importdata"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/issuetype"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/macro"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/priority"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/project"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/queue"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/report"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/resolution"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/sprint"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/status"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/trigger"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/user"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/version"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/workflow"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/worklog"
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
