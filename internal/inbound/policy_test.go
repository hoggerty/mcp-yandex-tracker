package inbound

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

type fakeTool struct{ name string }

func (f fakeTool) Description() mcp.Tool { return mcp.NewTool(f.name) }

func (f fakeTool) Handle(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return nil, nil
}

func entry(name string, level safety.Level) tool.Entry {
	return tool.Entry{Tool: fakeTool{name: name}, Safety: level}
}

func TestPolicyAllow(t *testing.T) {
	tests := []struct {
		name         string
		mode         safety.Mode
		enabledTools []string
		entry        tool.Entry
		want         bool
	}{
		{
			name:  "read-only mode allows read-only op",
			mode:  safety.ModeReadOnly,
			entry: entry("board_get", safety.ReadOnly),
			want:  true,
		},
		{
			name:  "read-only mode blocks write op",
			mode:  safety.ModeReadOnly,
			entry: entry("board_create", safety.Write),
			want:  false,
		},
		{
			name:  "full mode allows destructive op",
			mode:  safety.ModeFull,
			entry: entry("board_delete", safety.Destructive),
			want:  true,
		},
		{
			name:  "empty enabled tools allows everything",
			mode:  safety.ModeFull,
			entry: entry("board_get", safety.ReadOnly),
			want:  true,
		},
		{
			name:         "enabled wildcard pattern allows matching tool",
			mode:         safety.ModeFull,
			enabledTools: []string{"board_*"},
			entry:        entry("board_get", safety.ReadOnly),
			want:         true,
		},
		{
			name:         "enabled wildcard pattern does not allow non-matching tool",
			mode:         safety.ModeFull,
			enabledTools: []string{"board_*"},
			entry:        entry("queue_get", safety.ReadOnly),
			want:         false,
		},
		{
			name:         "enabled wildcard pattern does not bleed into a differently-prefixed domain",
			mode:         safety.ModeFull,
			enabledTools: []string{"entity_*"},
			entry:        entry("entitycomment_add", safety.Write),
			want:         false,
		},
		{
			name:         "enabled exact tool name allows just that tool",
			mode:         safety.ModeFull,
			enabledTools: []string{"board_get"},
			entry:        entry("board_get", safety.ReadOnly),
			want:         true,
		},
		{
			name:         "enabled exact tool name does not allow a different tool",
			mode:         safety.ModeFull,
			enabledTools: []string{"board_get"},
			entry:        entry("board_delete", safety.Destructive),
			want:         false,
		},
		{
			name:         "bare wildcard matches every tool",
			mode:         safety.ModeFull,
			enabledTools: []string{"*"},
			entry:        entry("queue_get", safety.ReadOnly),
			want:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Policy{Mode: tt.mode, EnabledTools: tt.enabledTools}
			if got := p.Allow(tt.entry); got != tt.want {
				t.Errorf("Allow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicyApply(t *testing.T) {
	entries := []tool.Entry{
		entry("board_get", safety.ReadOnly),
		entry("board_create", safety.Write),
		entry("board_delete", safety.Destructive),
		entry("queue_get", safety.ReadOnly),
	}

	tests := []struct {
		name         string
		mode         safety.Mode
		enabledTools []string
		want         []string
	}{
		{
			name: "read-only mode keeps only read-only ops",
			mode: safety.ModeReadOnly,
			want: []string{"board_get", "queue_get"},
		},
		{
			name: "read-write mode drops destructive op",
			mode: safety.ModeReadWrite,
			want: []string{"board_get", "board_create", "queue_get"},
		},
		{
			name: "full mode keeps everything",
			mode: safety.ModeFull,
			want: []string{"board_get", "board_create", "board_delete", "queue_get"},
		},
		{
			name:         "enabled wildcard pattern restricts to that domain",
			mode:         safety.ModeFull,
			enabledTools: []string{"board_*"},
			want:         []string{"board_get", "board_create", "board_delete"},
		},
		{
			name:         "enabled patterns combined restrict to the union",
			mode:         safety.ModeFull,
			enabledTools: []string{"board_get", "queue_*"},
			want:         []string{"board_get", "queue_get"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Policy{Mode: tt.mode, EnabledTools: tt.enabledTools}
			got := p.Apply(entries)

			names := make([]string, len(got))
			for i, e := range got {
				names[i] = e.Tool.Description().Name
			}
			if len(names) != len(tt.want) {
				t.Fatalf("Apply() = %v, want %v", names, tt.want)
			}
			for i := range names {
				if names[i] != tt.want[i] {
					t.Errorf("Apply()[%d] = %q, want %q", i, names[i], tt.want[i])
				}
			}
		})
	}
}
