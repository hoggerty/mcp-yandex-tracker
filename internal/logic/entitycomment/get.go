package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

// Get retrieves a single entity comment.
func (s *Service) Get(ctx context.Context, entityType, id, commentID string) (json.RawMessage, error) {
	return s.client.GetEntityComment(ctx, entityType, id, commentID)
}
