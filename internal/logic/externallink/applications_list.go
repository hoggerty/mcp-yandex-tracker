package externallink

import (
	"context"
	"encoding/json"
)

// ApplicationsList returns external applications known to Tracker (usable as link origins).
func (s *Service) ApplicationsList(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListApplications(ctx)
}
