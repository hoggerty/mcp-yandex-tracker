package project

import (
	"context"
	"encoding/json"
)

func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateProject(ctx, fields)
}
