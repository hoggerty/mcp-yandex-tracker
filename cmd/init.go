package cmd

import (
	"bufio"
	"os"

	"github.com/spf13/cobra"
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

func runInit() {
	in := bufio.NewScanner(os.Stdin)
	env := collectAndWriteEnv(in)
	dispatchInstall(in, env)
}
