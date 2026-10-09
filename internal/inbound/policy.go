package inbound

import (
	"path"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
	"github.com/hoggerty/yandex-tracker-mcp/internal/wire/config"
)

// Policy decides which registered tools are actually exposed to MCP
// clients. It is applied once, at startup, before AddTools — an entry that
// Policy excludes is absent from the server's tool list entirely, not
// merely hinted at as dangerous.
type Policy struct {
	Mode          config.ToolsMode
	DisabledTools []string
}

// Allow reports whether e should be registered under this policy.
func (p Policy) Allow(e tool.Entry) bool {
	if !modeAllows(p.Mode, e.Safety) {
		return false
	}
	name := e.Tool.Description().Name
	for _, pattern := range p.DisabledTools {
		if matched, _ := path.Match(pattern, name); matched {
			return false
		}
	}
	return true
}

// modeAllows applies the mode-level rule only, ignoring DisabledTools.
func modeAllows(mode config.ToolsMode, level safety.Level) bool {
	switch mode {
	case config.ToolsModeReadOnly:
		return level == safety.ReadOnly
	case config.ToolsModeFull:
		return true
	default: // config.ToolsModeReadWrite
		return level != safety.Destructive
	}
}

// Apply filters entries down to those allowed by the policy.
func (p Policy) Apply(entries []tool.Entry) []tool.Entry {
	allowed := make([]tool.Entry, 0, len(entries))
	for _, e := range entries {
		if p.Allow(e) {
			allowed = append(allowed, e)
		}
	}
	return allowed
}
