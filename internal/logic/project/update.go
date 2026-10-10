package project

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) UpdateSafety() safety.Level { return safety.Write }

func (s *Service) Update(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.UpdateSafety()); err != nil {
		return nil, err
	}

	return s.client.UpdateProject(ctx, id, version, fields)
}
