package user

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CurrentSafety = safety.ReadOnly

// Current retrieves the profile of the user that owns the API token.
func (s *Service) Current(ctx context.Context) (json.RawMessage, error) {
	return s.client.CurrentUser(ctx)
}
