package gap

import (
	"context"
	"encoding/json"
)

func (s *Service) Delete(ctx context.Context, gapIDs []string) (json.RawMessage, error) {
	return s.client.DeleteGaps(ctx, gapIDs)
}
