package resolution

import (
	"context"
	"encoding/json"
)

// Create creates a new issue resolution.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateResolution(ctx, fields)
}
