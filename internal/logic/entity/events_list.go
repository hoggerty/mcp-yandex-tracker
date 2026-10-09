package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const EventsListSafety = safety.ReadOnly

// EventsList retrieves an entity's change history, paginated relative to an event id.
func (s *Service) EventsList(ctx context.Context, entityType, id string, perPage int, from, selected string) (json.RawMessage, error) {
	return s.client.GetEntityEventsRelative(ctx, entityType, id, perPage, from, selected)
}
