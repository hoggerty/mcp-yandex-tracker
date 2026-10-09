package field

import (
	"context"
	"encoding/json"
)

// Get retrieves the parameters of a single global field.
func (s *Service) Get(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetField(ctx, id)
}
