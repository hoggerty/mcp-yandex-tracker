package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CreateSafety = safety.Write

// Create makes a new issue, merging queue/summary/description with any extra raw fields.
func (s *Service) Create(ctx context.Context, queue, summary, description string, extra map[string]any) (json.RawMessage, error) {
	fields := map[string]any{
		"queue":   map[string]any{"key": queue},
		"summary": summary,
	}
	if description != "" {
		fields["description"] = description
	}
	for k, v := range extra {
		fields[k] = v
	}
	return s.client.CreateIssue(ctx, fields)
}
