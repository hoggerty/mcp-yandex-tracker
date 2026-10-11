// Package actions contains the Yandex Tracker issue type use-case service.
package issuetype

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
	"github.com/hoggerty/mcp-yandex-tracker/internal/outbound/trackerapi"
)

type Service struct {
	client *trackerapi.Client
	guard  safety.Guard
}

func New(client *trackerapi.Client, guard safety.Guard) *Service {
	return &Service{client: client, guard: guard}
}
