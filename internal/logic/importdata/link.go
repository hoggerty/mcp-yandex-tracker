package importdata

import (
	"context"
	"encoding/json"
)

func (s *Service) Link(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return s.client.ImportLink(ctx, issueID, fields)
}
