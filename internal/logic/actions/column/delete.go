package column

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete deletes a board column; version must match the board's current version.
func (s *Service) Delete(ctx context.Context, boardID, columnID string, version int) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.DeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteColumn(ctx, boardID, columnID, version)
	})
}
