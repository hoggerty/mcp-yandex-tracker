package version

import (
	"context"
	"encoding/json"
)

// Get retrieves the parameters of a single version.
func (s *Service) Get(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetVersion(ctx, id)
}
