package board

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSprintsSafety = safety.ReadOnly

func (s *Service) ListSprints(ctx context.Context, boardID string) (json.RawMessage, error) {
	return s.client.ListBoardSprints(ctx, boardID)
}
