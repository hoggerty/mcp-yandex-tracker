// Package component contains the Yandex Tracker queue-component use-case service.
package component

import "github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"

type Service struct{ client *trackerapi.Client }

func New(client *trackerapi.Client) *Service { return &Service{client: client} }
