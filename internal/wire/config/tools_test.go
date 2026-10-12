package config

import (
	"testing"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func TestToolsConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     ToolsConfig
		wantErr bool
	}{
		{
			name: "default-shaped config is valid",
			cfg:  ToolsConfig{Mode: safety.ModeReadWrite, Tools: []string{"issue_*", "board*", "queue_get"}},
		},
		{
			name:    "invalid mode",
			cfg:     ToolsConfig{Mode: "bogus"},
			wantErr: true,
		},
		{
			name:    "malformed glob pattern",
			cfg:     ToolsConfig{Mode: safety.ModeReadWrite, Tools: []string{"issue_["}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
