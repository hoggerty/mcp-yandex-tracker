package board

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListSprintsSafety() safety.Level { return safety.ReadOnly }

func (s *Service) ListSprints(ctx context.Context, boardID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ListSprintsSafety()); err != nil {
		return nil, err
	}

	return s.client.ListBoardSprints(ctx, boardID)
}
