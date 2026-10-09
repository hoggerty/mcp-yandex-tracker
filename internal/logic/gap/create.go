package gap

import (
	"context"
	"encoding/json"
)

func (s *Service) Create(ctx context.Context, gaps []map[string]any) (json.RawMessage, error) {
	return s.client.CreateGaps(ctx, gaps)
}
