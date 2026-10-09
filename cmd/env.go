package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hoggerty/yandex-tracker-mcp/internal/wire/config"
)

type envVar struct {
	key   string
	value string
}

// requireEnv returns the currently configured credentials, prompting the
// user to run the init wizard if a required value is missing.
func requireEnv() ([]envVar, bool) {
	if env, ok := loadedEnv(); ok {
		return env, true
	}

	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Не заполнены обязательные переменные окружения: YANDEX_TRACKER_TOKEN и один из YANDEX_TRACKER_ORG_ID / YANDEX_TRACKER_CLOUD_ORG_ID.")
	if strings.ToLower(prompt(in, "Запустить init и заполнить их сейчас? [Y/n]", "y")) == "n" {
		return nil, false
	}

	return collectAndWriteEnv(in), true
}

// loadedEnv reads the Tracker config from the process environment (populated
// either directly or via .env, see cmd.go's init()), reusing the same
// parsing and "one of OrgID/CloudOrgID" validation as the server itself.
func loadedEnv() ([]envVar, bool) {
	cfg, err := config.Load()
	if err != nil {
		return nil, false
	}
	if err := cfg.TrackerClient.Validate(); err != nil {
		return nil, false
	}

	env := []envVar{{"YANDEX_TRACKER_TOKEN", cfg.TrackerClient.Token}}
	if cfg.TrackerClient.OrgID != "" {
		env = append(env, envVar{"YANDEX_TRACKER_ORG_ID", cfg.TrackerClient.OrgID})
	} else {
		env = append(env, envVar{"YANDEX_TRACKER_CLOUD_ORG_ID", cfg.TrackerClient.CloudOrgID})
	}
	// BaseURL always has a default (see TrackerClientConfig), so only pass it
	// along if the user actually overrode it.
	if baseURL := os.Getenv("YANDEX_TRACKER_BASE_URL"); baseURL != "" {
		env = append(env, envVar{"YANDEX_TRACKER_BASE_URL", baseURL})
	}

	return env, true
}
