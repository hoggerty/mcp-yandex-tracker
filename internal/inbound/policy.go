package inbound

import (
	"path"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

// Policy decides which registered tools are actually exposed to MCP
// clients. It is applied once, at startup, before AddTools — an entry that
// Policy excludes is absent from the server's tool list entirely, not
// merely hinted at as dangerous. Mode is the same type the logic layer uses
// to guard operation execution (safety.Guard), so a mode-disabled tool is
// unavailable at both layers from a single source of truth.
type Policy struct {
	Mode safety.Mode
	// EnabledTools lists glob patterns (as accepted by path.Match) matched
	// against a tool's name; a tool is exposed only if it matches at least
	// one of them. An empty slice disables this filter, exposing every tool
	// permitted by Mode.
	EnabledTools []string
}

// Allow reports whether e should be registered under this policy.
func (p Policy) Allow(e tool.Entry) bool {
	if !p.Mode.Allows(e.Safety) {
		return false
	}
	name := e.Tool.Description().Name
	return len(p.EnabledTools) == 0 || matchesAny(p.EnabledTools, name)
}

// matchesAny reports whether name matches any of the given glob patterns.
func matchesAny(patterns []string, name string) bool {
	for _, pattern := range patterns {
		if matched, _ := path.Match(pattern, name); matched {
			return true
		}
	}
	return false
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
