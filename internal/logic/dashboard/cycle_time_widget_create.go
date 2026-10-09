package dashboard

import (
	"context"
	"encoding/json"
)

// CreateCycleTimeWidget creates a "Cycle time" widget on a dashboard.
func (s *Service) CreateCycleTimeWidget(ctx context.Context, dashboardID string, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateCycleTimeWidget(ctx, dashboardID, fields)
}
