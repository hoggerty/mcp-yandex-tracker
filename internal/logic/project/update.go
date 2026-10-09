package project

import (
	"context"
	"encoding/json"
)

func (s *Service) Update(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.UpdateProject(ctx, id, version, fields)
}
