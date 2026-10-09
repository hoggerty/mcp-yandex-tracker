package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const TicketSafety = safety.Write

func (s *Service) Ticket(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.ImportTicket(ctx, fields)
}
