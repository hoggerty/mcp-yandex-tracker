package cmd

import (
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"

	"github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"
	"github.com/hoggerty/yandex-tracker-mcp/internal/wire"
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

	s := wire.NewServer(cfg, VERSION)
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "yandex-tracker-mcp: server error: %v\n", err)
		os.Exit(1)
	}
}
