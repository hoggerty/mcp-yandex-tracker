package board

import (
	"context"
	"encoding/json"
)

func (s *Service) ListSprints(ctx context.Context, boardID string) (json.RawMessage, error) {
	return s.client.ListBoardSprints(ctx, boardID)
}
