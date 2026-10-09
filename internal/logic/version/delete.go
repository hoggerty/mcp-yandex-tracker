package version

import (
	"context"
	"encoding/json"
)

// Delete deletes a version.
func (s *Service) Delete(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.DeleteVersion(ctx, id)
}
