package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) TicketSafety() safety.Level { return safety.Write }

func (s *Service) Ticket(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.TicketSafety()); err != nil {
		return nil, err
	}

	return s.client.ImportTicket(ctx, fields)
}
