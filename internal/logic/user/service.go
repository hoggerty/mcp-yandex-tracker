// Package user contains the Yandex Tracker user use-case service.
package user

import "github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"

type Service struct{ client *trackerapi.Client }

func New(client *trackerapi.Client) *Service { return &Service{client: client} }
