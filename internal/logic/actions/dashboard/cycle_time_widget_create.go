package dashboard

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CreateCycleTimeWidgetSafety() safety.Level { return safety.Write }

// CreateCycleTimeWidget creates a "Cycle time" widget on a dashboard.
func (s *Service) CreateCycleTimeWidget(ctx context.Context, dashboardID string, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.CreateCycleTimeWidgetSafety()); err != nil {
		return nil, err
	}

	return s.client.CreateCycleTimeWidget(ctx, dashboardID, fields)
}
