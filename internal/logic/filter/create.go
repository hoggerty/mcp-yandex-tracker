package filter

import (
	"context"
	"encoding/json"
)

// Create saves a new saved issue filter.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateFilter(ctx, fields)
}
