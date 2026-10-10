package column

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CreateSafety() safety.Level { return safety.Write }

// Create creates a new column on an agile board; version must match the board's current version.
func (s *Service) Create(ctx context.Context, boardID string, version int, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.CreateSafety(), func() (json.RawMessage, error) {
		return s.client.CreateColumn(ctx, boardID, version, fields)
	})
}
