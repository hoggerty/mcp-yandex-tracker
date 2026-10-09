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

func New(client *trackerapi.Client) *Services {
	return &Services{
		Issue:            issue.New(client),
		Queue:            queue.New(client),
		User:             user.New(client),
		Board:            board.New(client),
		Report:           report.New(client),
		Worklog:          worklog.New(client),
		ExternalLink:     externallink.New(client),
		BulkChange:       bulkchange.New(client),
		Field:            field.New(client),
		Filter:           filter.New(client),
		Workflow:         workflow.New(client),
		AutoAction:       autoaction.New(client),
		Trigger:          trigger.New(client),
		Macro:            macro.New(client),
		Component:        component.New(client),
		Version:          version.New(client),
		Column:           column.New(client),
		Sprint:           sprint.New(client),
		Dashboard:        dashboard.New(client),
		IssueType:        issuetype.New(client),
		Status:           status.New(client),
		Resolution:       resolution.New(client),
		Priority:         priority.New(client),
		Entity:           entity.New(client),
		EntityComment:    entitycomment.New(client),
		EntityChecklist:  entitychecklist.New(client),
		EntityAttachment: entityattachment.New(client),
		EntityLink:       entitylink.New(client),
		EntityAccess:     entityaccess.New(client),
		Project:          project.New(client),
		ImportData:       importdata.New(client),
		Gap:              gap.New(client),
	}
}
