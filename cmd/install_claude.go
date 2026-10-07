package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RunInstallClaude registers the server in Claude Code via `claude mcp add`.
var RunInstallClaude = &cobra.Command{
	Use:   "install-claude",
	Short: "register the server in Claude Code",
	Long:  "Register yandex-tracker-mcp as an MCP server in Claude Code (claude mcp add). If required environment variables are missing, offers to run init first.",
	Run: func(_ *cobra.Command, _ []string) {
		env, ok := requireEnv()
		if !ok {
			return
		}
		installClaude(env)
	},
}

func installClaude(env []envVar) {
	bin, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "yandex-tracker-mcp: %v\n", err)
		return
	}

	args := []string{"mcp", "add", "-s", "user"}
	for _, v := range env {
		args = append(args, "-e", v.key+"="+v.value)
	}
	args = append(args, "--transport", "stdio", "yandex-tracker", bin)

	runInstallCommand("claude", args)
}
