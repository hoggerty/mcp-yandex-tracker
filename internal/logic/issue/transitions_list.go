package issue

import (
	"context"
	"encoding/json"
)

func (s *Service) ListTransitions(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListTransitions(ctx, issueKey)
}
