package priority

import (
	"context"
	"encoding/json"
)

// List lists the issue priorities available in the organization. If
// localized is false, names are returned in all supported languages.
func (s *Service) List(ctx context.Context, localized bool) (json.RawMessage, error) {
	return s.client.ListPriorities(ctx, localized)
}
