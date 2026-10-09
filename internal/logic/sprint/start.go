package sprint

import (
	"context"
	"encoding/json"
)

// Start changes a sprint's status to "in progress". version is optional (0 to omit).
func (s *Service) Start(ctx context.Context, id string, version int) (json.RawMessage, error) {
	return s.client.StartSprint(ctx, id, version)
}
