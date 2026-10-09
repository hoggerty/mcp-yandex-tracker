package config

import "errors"

type TrackerClientConfig struct {
	Token      string `env:"YANDEX_TRACKER_TOKEN,required"`
	OrgID      string `env:"YANDEX_TRACKER_ORG_ID"`
	CloudOrgID string `env:"YANDEX_TRACKER_CLOUD_ORG_ID"`
	BaseURL    string `env:"YANDEX_TRACKER_BASE_URL" envDefault:"https://api.tracker.yandex.net"`
}

// Validate reports an error unless exactly one of OrgID or CloudOrgID is set.
func (c TrackerClientConfig) Validate() error {
	if c.OrgID == "" && c.CloudOrgID == "" {
		return errors.New("one of YANDEX_TRACKER_ORG_ID or YANDEX_TRACKER_CLOUD_ORG_ID is required")
	}
	return nil
}
