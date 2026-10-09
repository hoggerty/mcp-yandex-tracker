package entityattachment

import (
	"context"
	"encoding/json"
)

// Delete deletes a file attached to an entity.
func (s *Service) Delete(ctx context.Context, entityType, id, fileID string) (json.RawMessage, error) {
	return s.client.DeleteEntityAttachment(ctx, entityType, id, fileID)
}
