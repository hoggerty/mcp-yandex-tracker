package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete deletes an entity comment.
func (s *Service) Delete(ctx context.Context, entityType, id, commentID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.DeleteSafety()); err != nil {
		return nil, err
	}

	return s.client.DeleteEntityComment(ctx, entityType, id, commentID)
}
