package component

import (
	"context"
	"encoding/json"
)

// List lists all components in the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListComponents(ctx)
}
