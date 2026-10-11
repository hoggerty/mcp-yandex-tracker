package config

import (
	"fmt"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

type ToolsConfig struct {
	Mode safety.Mode `env:"YANDEX_TRACKER_TOOLS_MODE" envDefault:"read-write"`
	// Tools lists glob patterns (as accepted by path.Match) matched against
	// a tool's name; a tool is exposed only if it matches at least one of
	// them. The default covers the operations most Tracker workflows need
	// (issues, queues, boards, components, global fields, users, worklogs)
	// without the long tail of admin/bulk/entity tools that would otherwise
	// crowd a model's tool choice. Set to "*" to expose every tool permitted
	// by Mode.
	Tools []string `env:"YANDEX_TRACKER_TOOLS" envSeparator:"," envDefault:"issue_*,queue*,board*,component_*,field_*,user*,worklog_*"`
}

func (c ToolsConfig) Validate() error {
	if !c.Mode.Valid() {
		return fmt.Errorf("invalid YANDEX_TRACKER_TOOLS_MODE %q: must be one of %q, %q, %q",
			c.Mode, safety.ModeReadOnly, safety.ModeReadWrite, safety.ModeFull)
	}
	return nil
}
