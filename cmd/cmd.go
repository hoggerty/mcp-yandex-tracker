// Package cmd defines the yandex-tracker-mcp CLI commands.
package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const defaultVersion = "dev"

// VERSION is the yandex-tracker-mcp version, set by -ldflags during build (see Makefile)
var VERSION = defaultVersion

var rootCmd = &cobra.Command{
	Use:   "yandex-tracker-mcp",
	Short: "yandex-tracker-mcp",
	Long:  "An MCP server that exposes Yandex Tracker operations as tools for MCP-capable clients.",
}

func init() {
	// .env is optional
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "yandex-tracker-mcp: failed to load .env: %v\n", err)
		os.Exit(1)
	}

	rootCmd.AddCommand(RunServer)
	rootCmd.AddCommand(RunInit)
	rootCmd.AddCommand(RunInstallClaude)
	rootCmd.AddCommand(RunInstallCodex)
	rootCmd.AddCommand(RunSnippet)
	rootCmd.Version = VERSION
}

func Execute() {
	if len(os.Args[1:]) == 0 {
		os.Args = append(os.Args, RunServer.Use)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
