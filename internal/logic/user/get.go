package user

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

// Get retrieves a single user by login or numeric id.
func (s *Service) Get(ctx context.Context, idOrLogin string) (json.RawMessage, error) {
	return s.client.GetUser(ctx, idOrLogin)
}
