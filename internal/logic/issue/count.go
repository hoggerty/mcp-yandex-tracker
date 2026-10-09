package issue

import (
	"context"
	"encoding/json"
)

func (s *Service) Count(ctx context.Context, query string, filter map[string]any) (json.RawMessage, error) {
	return s.client.CountIssues(ctx, query, filter)
}
