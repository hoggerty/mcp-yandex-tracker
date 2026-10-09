package externallink

import (
	"context"
	"encoding/json"
)

// Add links an issue to an object in an external application.
func (s *Service) Add(ctx context.Context, issueKey, relationship, key, origin string) (json.RawMessage, error) {
	return s.client.AddExternalLink(ctx, issueKey, relationship, key, origin)
}
