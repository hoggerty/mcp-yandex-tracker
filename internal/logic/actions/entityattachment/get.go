package entityattachment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) GetSafety() safety.Level { return safety.ReadOnly }

// Get retrieves metadata of a single file attached to an entity.
func (s *Service) Get(ctx context.Context, entityType, id, fileID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.GetSafety(), func() (json.RawMessage, error) {
		return s.client.GetEntityAttachment(ctx, entityType, id, fileID)
	})
}
