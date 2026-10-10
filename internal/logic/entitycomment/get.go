package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) GetSafety() safety.Level { return safety.ReadOnly }

// Get retrieves a single entity comment.
func (s *Service) Get(ctx context.Context, entityType, id, commentID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.GetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetEntityComment(ctx, entityType, id, commentID)
}
