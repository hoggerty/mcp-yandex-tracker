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
	Mode          safety.Mode
	DisabledTools []string
}

// Allow reports whether e should be registered under this policy.
func (p Policy) Allow(e tool.Entry) bool {
	if !p.Mode.Allows(e.Safety) {
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
