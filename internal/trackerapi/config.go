package trackerapi

import (
	"errors"
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Token      string `env:"YANDEX_TRACKER_TOKEN,required"`
	OrgID      string `env:"YANDEX_TRACKER_ORG_ID"`
	CloudOrgID string `env:"YANDEX_TRACKER_CLOUD_ORG_ID"`
	BaseURL    string `env:"YANDEX_TRACKER_BASE_URL" envDefault:"https://api.tracker.yandex.net"`
}

func ConfigFromEnv() (Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	if cfg.OrgID == "" && cfg.CloudOrgID == "" {
		return Config{}, errors.New("one of YANDEX_TRACKER_ORG_ID or YANDEX_TRACKER_CLOUD_ORG_ID is required")
	}

	return cfg, nil
}
