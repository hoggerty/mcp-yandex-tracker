// Package entityattachment contains the Yandex Tracker entity attachment use-case service.
package entityattachment

import "github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"

type Service struct{ client *trackerapi.Client }

func New(client *trackerapi.Client) *Service { return &Service{client: client} }
