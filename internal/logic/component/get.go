package component

import (
	"context"
	"encoding/json"
)

// Get retrieves the parameters of a single component.
func (s *Service) Get(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetComponent(ctx, id)
}
