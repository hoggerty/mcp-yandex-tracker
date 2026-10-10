// Package actions contains the Yandex Tracker entity checklist use-case service.
package entitychecklist

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
	"github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"
)

type Service struct {
	client *trackerapi.Client
	guard  safety.Guard
}

func New(client *trackerapi.Client, guard safety.Guard) *Service {
	return &Service{client: client, guard: guard}
}
