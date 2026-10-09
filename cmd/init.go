package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"
)

// RunInit is the interactive setup wizard: it writes a .env file and,
// optionally, registers the server with an MCP client (Claude Code, Codex).
var RunInit = &cobra.Command{
	Use:   "init",
	Short: "interactive setup wizard",
	Long:  "Interactively configure yandex-tracker-mcp: writes a .env file and optionally registers the server with an MCP client (Claude Code, Codex).",
	Run: func(_ *cobra.Command, _ []string) {
		runInit()
	},
}

type envVar struct {
	key   string
	value string
}

func runInit() {
	in := bufio.NewScanner(os.Stdin)
	env := collectAndWriteEnv(in)
	dispatchInstall(in, env)
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
	cfg, err := trackerapi.ConfigFromEnv()
	if err != nil {
		return nil, false
	}

	env := []envVar{{"YANDEX_TRACKER_TOKEN", cfg.Token}}
	if cfg.OrgID != "" {
		env = append(env, envVar{"YANDEX_TRACKER_ORG_ID", cfg.OrgID})
	} else {
		env = append(env, envVar{"YANDEX_TRACKER_CLOUD_ORG_ID", cfg.CloudOrgID})
	}
	// BaseURL always has a default (see Config), so only pass it along if the
	// user actually overrode it.
	if baseURL := os.Getenv("YANDEX_TRACKER_BASE_URL"); baseURL != "" {
		env = append(env, envVar{"YANDEX_TRACKER_BASE_URL", baseURL})
	}

	return env, true
}

func collectAndWriteEnv(in *bufio.Scanner) []envVar {
	token := prompt(in, "OAuth-токен Яндекс Трекера (YANDEX_TRACKER_TOKEN)", "")
	for token == "" {
		fmt.Fprintln(os.Stderr, "Токен обязателен.")
		token = prompt(in, "OAuth-токен Яндекс Трекера (YANDEX_TRACKER_TOKEN)", "")
	}

	orgKey, orgValue := promptOrg(in)
	baseURL := prompt(in, "Базовый URL API (YANDEX_TRACKER_BASE_URL, Enter — значение по умолчанию)", "")

	env := []envVar{
		{"YANDEX_TRACKER_TOKEN", token},
		{orgKey, orgValue},
	}
	if baseURL != "" {
		env = append(env, envVar{"YANDEX_TRACKER_BASE_URL", baseURL})
	}

	if err := writeEnvFile(in, ".env", env); err != nil {
		fmt.Fprintf(os.Stderr, "yandex-tracker-mcp: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Записано в .env")

	return env
}

// dispatchInstall asks whether to register the server with an agent and
// runs the matching installer, shared by init and the standalone commands.
func dispatchInstall(in *bufio.Scanner, env []envVar) {
	switch strings.ToLower(prompt(in, "Установить сервер для агента? [claude/codex/other/нет]", "нет")) {
	case "claude":
		installClaude(env)
	case "codex":
		installCodex(env)
	case "other", "другое", "другие":
		printSnippet(env)
	}
}

func promptOrg(in *bufio.Scanner) (key, value string) {
	for {
		choice := prompt(in, "Тип организации: 1) Яндекс 360 (X-Org-ID)  2) Yandex Cloud (X-Cloud-Org-ID)", "1")
		switch choice {
		case "1", "":
			id := prompt(in, "ID организации (YANDEX_TRACKER_ORG_ID)", "")
			if id == "" {
				fmt.Fprintln(os.Stderr, "ID организации обязателен.")
				continue
			}
			return "YANDEX_TRACKER_ORG_ID", id
		case "2":
			id := prompt(in, "ID организации (YANDEX_TRACKER_CLOUD_ORG_ID)", "")
			if id == "" {
				fmt.Fprintln(os.Stderr, "ID организации обязателен.")
				continue
			}
			return "YANDEX_TRACKER_CLOUD_ORG_ID", id
		default:
			fmt.Fprintln(os.Stderr, "Введите 1 или 2.")
		}
	}
}

func prompt(in *bufio.Scanner, question, def string) string {
	if def != "" {
		fmt.Printf("%s [%s]: ", question, def)
	} else {
		fmt.Printf("%s: ", question)
	}

	if !in.Scan() {
		return def
	}

	line := strings.TrimSpace(in.Text())
	if line == "" {
		return def
	}

	return line
}

func writeEnvFile(in *bufio.Scanner, path string, env []envVar) error {
	if _, err := os.Stat(path); err == nil {
		if strings.ToLower(prompt(in, fmt.Sprintf("Файл %s уже существует. Перезаписать? [y/N]", path), "n")) != "y" {
			fmt.Println("Отменено.")
			os.Exit(0)
		}
	}

	var b strings.Builder
	for _, v := range env {
		fmt.Fprintf(&b, "%s=%s\n", v.key, v.value)
	}

	return os.WriteFile(path, []byte(b.String()), 0o600)
}

func runInstallCommand(name string, args []string) {
	c := exec.Command(name, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "yandex-tracker-mcp: %s %s: %v\n", name, strings.Join(args, " "), err)
	}
}
