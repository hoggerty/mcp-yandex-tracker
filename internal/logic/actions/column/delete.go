package column

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete deletes a board column; version must match the board's current version.
func (s *Service) Delete(ctx context.Context, boardID, columnID string, version int) (json.RawMessage, error) {
	if err := s.guard.Check(s.DeleteSafety()); err != nil {
		return nil, err
	}

	return s.client.DeleteColumn(ctx, boardID, columnID, version)
}
