package externallink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ApplicationsListSafety() safety.Level { return safety.ReadOnly }

// ApplicationsList returns external applications known to Tracker (usable as link origins).
func (s *Service) ApplicationsList(ctx context.Context) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ApplicationsListSafety(), func() (json.RawMessage, error) {
		return s.client.ListApplications(ctx)
	})
}
