package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"

	"github.com/hoggerty/mcp-yandex-tracker/internal/wire"
)

const (
	transportStdio = "stdio"
	transportHTTP  = "http"
)

var (
	transport string
	httpAddr  string
)

// RunServer starts the MCP server over stdio or Streamable HTTP.
var RunServer = &cobra.Command{
	Use:   "server",
	Short: "start the MCP server",
	Long:  "Start the yandex-tracker-mcp server, communicating over stdio or Streamable HTTP.",
	Run: func(_ *cobra.Command, _ []string) {
		runServer()
	},
}

func init() {
	RunServer.Flags().StringVar(&transport, "transport", transportStdio, `transport to serve on: "stdio" or "http"`)
	RunServer.Flags().StringVar(&httpAddr, "http-addr", ":8080", `address to listen on when --transport=http`)
}

func runServer() {
	s, err := wire.NewServer(VERSION)
	if err != nil {
		fatalf("%v", err)
	}

	switch transport {
	case transportStdio:
		if err := server.ServeStdio(s); err != nil {
			fatalf("server error: %v", err)
		}
	case transportHTTP:
		serveHTTP(s)
	default:
		fatalf(`unknown --transport %q: expected "stdio" or "http"`, transport)
	}
}

// serveHTTP runs the Streamable HTTP transport until it fails or the
// process receives an interrupt/termination signal, in which case it shuts
// down gracefully.
func serveHTTP(s *server.MCPServer) {
	httpServer := server.NewStreamableHTTPServer(s)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		errCh <- httpServer.Start(httpAddr)
	}()

	select {
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			fatalf("server error: %v", err)
		}
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fatalf("server shutdown error: %v", err)
		}
	}
}
