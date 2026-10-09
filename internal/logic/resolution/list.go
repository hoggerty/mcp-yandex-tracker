package resolution

import (
	"context"
	"encoding/json"
)

// List lists the issue resolutions available in the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListResolutions(ctx)
}
