// Command yandex-tracker-mcp runs an MCP server that exposes Yandex Tracker
// operations (issues, queues, boards, comments, transitions, users) as tools
// for MCP-capable clients, communicating over stdio.
package main

import (
	"github.com/hoggerty/yandex-tracker-mcp/cmd"
)

func main() {
	cmd.Execute()
}
