package project

import (
	"context"
	"encoding/json"
)

func (s *Service) Get(ctx context.Context, id, expand string) (json.RawMessage, error) {
	return s.client.GetProject(ctx, id, expand)
}
