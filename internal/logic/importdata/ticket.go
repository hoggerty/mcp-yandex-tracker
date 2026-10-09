package importdata

import (
	"context"
	"encoding/json"
)

func (s *Service) Ticket(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.ImportTicket(ctx, fields)
}
