package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RunInstallCodex registers the server in Codex via `codex mcp add`.
var RunInstallCodex = &cobra.Command{
	Use:   "install-codex",
	Short: "register the server in Codex",
	Long:  "Register yandex-tracker-mcp as an MCP server in Codex (codex mcp add). If required environment variables are missing, offers to run init first.",
	Run: func(_ *cobra.Command, _ []string) {
		env, ok := requireEnv()
		if !ok {
			return
		}
		installCodex(env)
	},
}

func installCodex(env []envVar) {
	bin, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "yandex-tracker-mcp: %v\n", err)
		return
	}

	args := []string{"mcp", "add", "yandex-tracker"}
	for _, v := range env {
		args = append(args, "--env", v.key+"="+v.value)
	}
	args = append(args, "--", bin, "server")

	runInstallCommand("codex", args)
}
