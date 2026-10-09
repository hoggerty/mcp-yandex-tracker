package field

import (
	"context"
	"encoding/json"
)

// List returns all global fields of the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListFields(ctx)
}
