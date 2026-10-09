package dashboard

import (
	"context"
	"encoding/json"
)

// Create creates a new dashboard.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateDashboard(ctx, fields)
}
