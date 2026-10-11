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
		name          string
		mode          safety.Mode
		disabledTools []string
		entry         tool.Entry
		want          bool
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
			name:          "disabled glob blocks matching name",
			mode:          safety.ModeFull,
			disabledTools: []string{"board_*"},
			entry:         entry("board_get", safety.ReadOnly),
			want:          false,
		},
		{
			name:          "disabled glob does not block non-matching name",
			mode:          safety.ModeFull,
			disabledTools: []string{"board_*"},
			entry:         entry("queue_get", safety.ReadOnly),
			want:          true,
		},
		{
			name:          "disabled exact name blocks that tool",
			mode:          safety.ModeFull,
			disabledTools: []string{"queue_delete"},
			entry:         entry("queue_delete", safety.Destructive),
			want:          false,
		},
		{
			name:          "mode allows but disabled list still blocks",
			mode:          safety.ModeReadOnly,
			disabledTools: []string{"*"},
			entry:         entry("board_get", safety.ReadOnly),
			want:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Policy{Mode: tt.mode, DisabledTools: tt.disabledTools}
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
		name          string
		mode          safety.Mode
		disabledTools []string
		want          []string
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
			name:          "full mode minus disabled glob",
			mode:          safety.ModeFull,
			disabledTools: []string{"board_*"},
			want:          []string{"queue_get"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Policy{Mode: tt.mode, DisabledTools: tt.disabledTools}
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
