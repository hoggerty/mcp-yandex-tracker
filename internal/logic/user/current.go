package user

import (
	"context"
	"encoding/json"
)

// Current retrieves the profile of the user that owns the API token.
func (s *Service) Current(ctx context.Context) (json.RawMessage, error) {
	return s.client.CurrentUser(ctx)
}
