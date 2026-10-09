package component

import (
	"context"
	"encoding/json"
)

// Delete deletes a component.
func (s *Service) Delete(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.DeleteComponent(ctx, id)
}
