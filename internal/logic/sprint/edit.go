package sprint

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) EditSafety() safety.Level { return safety.Write }

// Edit edits a sprint; version must match the sprint's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.EditSafety()); err != nil {
		return nil, err
	}

	return s.client.EditSprint(ctx, id, version, fields)
}
