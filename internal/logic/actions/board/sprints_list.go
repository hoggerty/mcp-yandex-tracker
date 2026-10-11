package board

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ListSprintsSafety() safety.Level { return safety.ReadOnly }

func (s *Service) ListSprints(ctx context.Context, boardID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ListSprintsSafety(), func() (json.RawMessage, error) {
		return s.client.ListBoardSprints(ctx, boardID)
	})
}
