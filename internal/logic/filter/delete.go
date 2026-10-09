package filter

import (
	"context"
	"encoding/json"
)

// Delete deletes a saved issue filter.
func (s *Service) Delete(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.DeleteFilter(ctx, id)
}
