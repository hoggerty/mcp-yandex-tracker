package cmd

import (
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"

	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools"
	"github.com/hoggerty/yandex-tracker-mcp/internal/trackerapi"
)

// RunServer starts the MCP server over stdio.
var RunServer = &cobra.Command{
	Use:   "server",
	Short: "start the MCP server",
	Long:  "Start the yandex-tracker-mcp server, communicating over stdio.",
	Run: func(_ *cobra.Command, _ []string) {
		runServer()
	},
}

func runServer() {
	cfg, err := trackerapi.ConfigFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "yandex-tracker-mcp: %v\n", err)
		os.Exit(1)
	}

	s := server.NewMCPServer("yandex-tracker-mcp", VERSION)
	mcptools.RegisterTools(s, trackerapi.New(cfg))

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "yandex-tracker-mcp: server error: %v\n", err)
		os.Exit(1)
	}
}
