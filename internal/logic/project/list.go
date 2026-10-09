package project

import (
	"context"
	"encoding/json"
)

func (s *Service) List(ctx context.Context, expand string) (json.RawMessage, error) {
	return s.client.ListProjects(ctx, expand)
}
