package entitylink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const DeleteSafety = safety.Destructive

// Delete removes the link between an entity and the entity identified by rightID.
func (s *Service) Delete(ctx context.Context, entityType, id, rightID string) (json.RawMessage, error) {
	return s.client.DeleteEntityLink(ctx, entityType, id, rightID)
}
