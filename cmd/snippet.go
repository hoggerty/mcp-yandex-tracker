package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// RunSnippet prints an mcpServers JSON snippet for clients without a
// dedicated installer (e.g. Cursor, Windsurf, Docker MCP Hub).
var RunSnippet = &cobra.Command{
	Use:   "snippet",
	Short: "print an mcpServers JSON snippet",
	Long:  "Print an mcpServers JSON snippet for MCP clients that don't have a dedicated installer. If required environment variables are missing, offers to run init first.",
	Run: func(_ *cobra.Command, _ []string) {
		env, ok := requireEnv()
		if !ok {
			return
		}
		printSnippet(env)
	},
}

type mcpServerEntry struct {
	Command string            `json:"command"`
	Env     map[string]string `json:"env"`
}

type mcpServersSnippet struct {
	McpServers map[string]mcpServerEntry `json:"mcpServers"`
}

func printSnippet(env []envVar) {
	bin := currentBinary()

	envMap := make(map[string]string, len(env))
	for _, v := range env {
		envMap[v.key] = v.value
	}

	snippet := mcpServersSnippet{
		McpServers: map[string]mcpServerEntry{
			"yandex-tracker": {Command: bin, Env: envMap},
		},
	}

	out, err := json.MarshalIndent(snippet, "", "  ")
	if err != nil {
		fatalf("%v", err)
	}

	fmt.Println(string(out))
}
