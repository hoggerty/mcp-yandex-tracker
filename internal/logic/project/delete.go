package project

import (
	"context"
	"encoding/json"
)

func (s *Service) Delete(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.DeleteProject(ctx, id)
}
