package entityattachment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete deletes a file attached to an entity.
func (s *Service) Delete(ctx context.Context, entityType, id, fileID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.DeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteEntityAttachment(ctx, entityType, id, fileID)
	})
}
