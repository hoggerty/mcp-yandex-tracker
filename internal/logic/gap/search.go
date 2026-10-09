package gap

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const SearchSafety = safety.ReadOnly

func (s *Service) Search(ctx context.Context, body map[string]any, perPage, page int) (json.RawMessage, error) {
	return s.client.SearchGaps(ctx, body, perPage, page)
}
