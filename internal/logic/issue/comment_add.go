package issue

import (
	"context"
	"encoding/json"
)

func (s *Service) AddComment(ctx context.Context, issueKey, text string, summonees []string) (json.RawMessage, error) {
	return s.client.AddComment(ctx, issueKey, text, summonees)
}
