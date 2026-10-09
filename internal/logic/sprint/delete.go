package sprint

import (
	"context"
	"encoding/json"
)

// Delete deletes a sprint.
func (s *Service) Delete(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.DeleteSprint(ctx, id)
}
