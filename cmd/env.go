package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"

	"github.com/hoggerty/mcp-yandex-tracker/internal/wire/config"
)

type envVar struct {
	key   string
	value string
}

// loadEnvFile loads .env into the process environment. It looks in the
// current directory first, matching where `init` writes it when run from a
// repo checkout (see writeEnvFile in prompt.go). If that's not found — e.g.
// the binary was installed elsewhere and is invoked by name via PATH, or run
// by something like systemd with a different working directory — it falls
// back to a .env next to the executable itself, so a copy of the binary
// with its own .env alongside it keeps working regardless of cwd.
//
// Neither call overrides variables already set in the real environment
// (godotenv.Load's behavior), so this never shadows explicit env vars.
func loadEnvFile() {
	if err := godotenv.Load(); err == nil {
		return
	} else if !os.IsNotExist(err) {
		fatalf("failed to load .env: %v", err)
	}

	bin, err := os.Executable()
	if err != nil {
		return
	}

	envPath := filepath.Join(filepath.Dir(bin), ".env")
	if err := godotenv.Load(envPath); err != nil && !os.IsNotExist(err) {
		fatalf("failed to load %s: %v", envPath, err)
	}
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
// either directly or via .env, see loadEnvFile), reusing the same parsing
// and "one of OrgID/CloudOrgID" validation as the server itself.
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
