package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const AttachmentListSafety = safety.ReadOnly

// AttachmentList lists the files attached to an issue.
func (s *Service) AttachmentList(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListAttachments(ctx, issueKey)
}
