package user

import (
	"context"
	"encoding/json"
)

// List lists users known to the organization.
func (s *Service) List(ctx context.Context, perPage, page int) (json.RawMessage, error) {
	return s.client.ListUsers(ctx, perPage, page)
}
