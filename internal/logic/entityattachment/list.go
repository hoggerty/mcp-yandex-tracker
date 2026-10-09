package entityattachment

import (
	"context"
	"encoding/json"
)

// List lists the files attached to an entity.
func (s *Service) List(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	return s.client.ListEntityAttachments(ctx, entityType, id)
}
