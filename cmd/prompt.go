package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		fatalf("%v", err)
	}
	fmt.Println("Записано в .env")

	return env
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
		fmt.Println()
		fatalf("стандартный ввод закрыт — интерактивная настройка невозможна. Задайте переменные окружения напрямую или через .env и повторите запуск.")
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
