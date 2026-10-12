package config

import (
	"errors"
	"fmt"
	"path"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

type ToolsConfig struct {
	Mode safety.Mode `env:"YANDEX_TRACKER_TOOLS_MODE" envDefault:"read-write"`
	// Tools lists glob patterns (as accepted by path.Match) matched against
	// a tool's name; a tool is exposed only if it matches at least one of
	// them. The default covers the operations most Tracker workflows need —
	// issues, read-only queue/component/field reference data, boards,
	// sprints, users, worklogs — without admin (create/edit/delete) and
	// long-tail/bulk/entity tools that would otherwise crowd a model's tool
	// choice. Set to "*" to expose every tool permitted by Mode — an empty
	// value does NOT do this: caarlos0/env falls back to envDefault for an
	// empty env var, so YANDEX_TRACKER_TOOLS= just re-applies the default
	// below instead of clearing it.
	Tools []string `env:"YANDEX_TRACKER_TOOLS" envSeparator:"," envDefault:"issue_*,queue_get,queues_list,queue_fields_list,queue_tags_list,queue_access_user_get,queue_access_group_get,queue_local_fields_list,queue_local_field_get,board*,sprint_*,component_get,component_list,component_queue_list,component_access_user_get,component_access_group_get,field_get,field_list,user*,worklog_*"`
}

func (c ToolsConfig) Validate() error {
	if !c.Mode.Valid() {
		return fmt.Errorf("invalid YANDEX_TRACKER_TOOLS_MODE %q: must be one of %q, %q, %q",
			c.Mode, safety.ModeReadOnly, safety.ModeReadWrite, safety.ModeFull)
	}
	for _, toolPath := range c.Tools {
		if _, err := path.Match(toolPath, ""); errors.Is(err, path.ErrBadPattern) {
			return fmt.Errorf("invalid YANDEX_TRACKER_TOOLS pattern %q: %w", toolPath, err)
		}
	}
	return nil
}
