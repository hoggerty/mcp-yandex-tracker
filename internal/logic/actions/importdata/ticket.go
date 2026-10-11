package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) TicketSafety() safety.Level { return safety.Write }

func (s *Service) Ticket(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.TicketSafety(), func() (json.RawMessage, error) {
		return s.client.ImportTicket(ctx, fields)
	})
}
