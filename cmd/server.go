package cmd

import (
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"

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
	s, err := wire.NewServer(VERSION)
	if err != nil {
		fatalf("%v", err)
	}

	if err := server.ServeStdio(s); err != nil {
		fatalf("server error: %v", err)
	}
}
