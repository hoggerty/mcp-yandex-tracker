package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const AddSafety = safety.Write

// Add adds a comment to an entity.
func (s *Service) Add(ctx context.Context, entityType, id string, fields map[string]any) (json.RawMessage, error) {
	return s.client.AddEntityComment(ctx, entityType, id, fields)
}
