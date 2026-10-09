package filter

import (
	"context"
	"encoding/json"
)

// Get retrieves a saved issue filter by id.
func (s *Service) Get(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetFilter(ctx, id)
}
