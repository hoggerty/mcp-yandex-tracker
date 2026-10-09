// Package entitychecklist contains the Yandex Tracker entity checklist use-case service.
package entitychecklist

import "github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"

type Service struct{ client *trackerapi.Client }

func New(client *trackerapi.Client) *Service { return &Service{client: client} }
