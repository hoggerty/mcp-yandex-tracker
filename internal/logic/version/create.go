package version

import (
	"context"
	"encoding/json"
)

// Create creates a new version on a queue.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateVersion(ctx, fields)
}
