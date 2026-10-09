package sprint

import (
	"context"
	"encoding/json"
)

// Get retrieves the parameters of a single sprint.
func (s *Service) Get(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetSprint(ctx, id)
}
