package project

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CreateSafety() safety.Level { return safety.Write }

func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.CreateSafety()); err != nil {
		return nil, err
	}

	return s.client.CreateProject(ctx, fields)
}
