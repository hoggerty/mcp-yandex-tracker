package config

import (
	"fmt"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

type ToolsConfig struct {
	Mode          safety.Mode `env:"YANDEX_TRACKER_TOOLS_MODE" envDefault:"read-write"`
	DisabledTools []string    `env:"YANDEX_TRACKER_DISABLED_TOOLS" envSeparator:","`
}

func (c ToolsConfig) Validate() error {
	if !c.Mode.Valid() {
		return fmt.Errorf("invalid YANDEX_TRACKER_TOOLS_MODE %q: must be one of %q, %q, %q",
			c.Mode, safety.ModeReadOnly, safety.ModeReadWrite, safety.ModeFull)
	}
	return nil
}
