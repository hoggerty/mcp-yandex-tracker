package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AttachmentGetSafety() safety.Level { return safety.ReadOnly }

// AttachmentGet retrieves metadata for a single attached file.
func (s *Service) AttachmentGet(ctx context.Context, issueKey, attachmentID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.AttachmentGetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetAttachmentInfo(ctx, issueKey, attachmentID)
}
