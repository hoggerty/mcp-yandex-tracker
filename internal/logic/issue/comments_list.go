package issue

import (
	"context"
	"encoding/json"
)

func (s *Service) ListComments(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListComments(ctx, issueKey)
}
