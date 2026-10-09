package config

import "fmt"

type ToolsMode string

const (
	ToolsModeReadOnly  ToolsMode = "read-only"
	ToolsModeReadWrite ToolsMode = "read-write"
	ToolsModeFull      ToolsMode = "full"
)

type ToolsConfig struct {
	Mode          ToolsMode `env:"YANDEX_TRACKER_TOOLS_MODE" envDefault:"read-write"`
	DisabledTools []string  `env:"YANDEX_TRACKER_DISABLED_TOOLS" envSeparator:","`
}

func (c ToolsConfig) Validate() error {
	switch c.Mode {
	case ToolsModeReadOnly, ToolsModeReadWrite, ToolsModeFull:
		return nil
	default:
		return fmt.Errorf("invalid YANDEX_TRACKER_TOOLS_MODE %q: must be one of %q, %q, %q",
			c.Mode, ToolsModeReadOnly, ToolsModeReadWrite, ToolsModeFull)
	}
}
