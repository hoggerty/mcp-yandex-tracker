package project

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) GetSafety() safety.Level { return safety.ReadOnly }

func (s *Service) Get(ctx context.Context, id, expand string) (json.RawMessage, error) {
	if err := s.guard.Check(s.GetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetProject(ctx, id, expand)
}
