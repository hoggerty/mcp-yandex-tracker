package macro

import (
	"context"
	"encoding/json"
)

// Create creates a new macro on a queue.
func (s *Service) Create(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateMacro(ctx, queueKey, fields)
}
