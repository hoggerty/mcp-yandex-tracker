package field

import (
	"context"
	"encoding/json"
)

// Create creates a new global issue field.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateField(ctx, fields)
}
