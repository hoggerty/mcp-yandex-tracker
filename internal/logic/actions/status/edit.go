package status

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) EditSafety() safety.Level { return safety.Write }

// Edit edits an issue status; version must match the status's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.EditSafety(), func() (json.RawMessage, error) {
		return s.client.EditStatus(ctx, id, version, fields)
	})
}
