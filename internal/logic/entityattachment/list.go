package entityattachment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

// List lists the files attached to an entity.
func (s *Service) List(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	return s.client.ListEntityAttachments(ctx, entityType, id)
}
