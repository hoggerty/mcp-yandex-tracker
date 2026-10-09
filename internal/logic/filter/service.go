// Package filter contains the Yandex Tracker saved-filter use-case service.
package filter

import "github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"

type Service struct{ client *trackerapi.Client }

func New(client *trackerapi.Client) *Service { return &Service{client: client} }
