package entitycomment

import (
	"context"
	"encoding/json"
)

// List lists the comments of an entity. expand is optional (e.g. "all", "html", "attachments", "reactions").
func (s *Service) List(ctx context.Context, entityType, id, expand string) (json.RawMessage, error) {
	return s.client.ListEntityComments(ctx, entityType, id, expand)
}
