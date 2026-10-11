package safety

import (
	"errors"
	"testing"
)

func TestModeValid(t *testing.T) {
	tests := []struct {
		name string
		mode Mode
		want bool
	}{
		{"read-only", ModeReadOnly, true},
		{"read-write", ModeReadWrite, true},
		{"full", ModeFull, true},
		{"empty", Mode(""), false},
		{"unknown", Mode("admin"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mode.Valid(); got != tt.want {
				t.Errorf("Mode(%q).Valid() = %v, want %v", tt.mode, got, tt.want)
			}
		})
	}
}

func TestModeAllows(t *testing.T) {
	tests := []struct {
		mode  Mode
		level Level
		want  bool
	}{
		{ModeReadOnly, ReadOnly, true},
		{ModeReadOnly, Write, false},
		{ModeReadOnly, Destructive, false},

		{ModeReadWrite, ReadOnly, true},
		{ModeReadWrite, Write, true},
		{ModeReadWrite, Destructive, false},

		{ModeFull, ReadOnly, true},
		{ModeFull, Write, true},
		{ModeFull, Destructive, true},
	}
	for _, tt := range tests {
		t.Run(string(tt.mode)+"/"+levelName(tt.level), func(t *testing.T) {
			if got := tt.mode.Allows(tt.level); got != tt.want {
				t.Errorf("Mode(%q).Allows(%s) = %v, want %v", tt.mode, levelName(tt.level), got, tt.want)
			}
		})
	}
}

func TestGuardCheck(t *testing.T) {
	tests := []struct {
		mode    Mode
		level   Level
		wantErr error
	}{
		{ModeReadOnly, ReadOnly, nil},
		{ModeReadOnly, Write, ErrDisabled},
		{ModeReadOnly, Destructive, ErrDisabled},

		{ModeReadWrite, Write, nil},
		{ModeReadWrite, Destructive, ErrDisabled},

		{ModeFull, Destructive, nil},
	}
	for _, tt := range tests {
		t.Run(string(tt.mode)+"/"+levelName(tt.level), func(t *testing.T) {
			g := NewGuard(tt.mode)
			if err := g.Check(tt.level); err != tt.wantErr {
				t.Errorf("Guard(%q).Check(%s) = %v, want %v", tt.mode, levelName(tt.level), err, tt.wantErr)
			}
		})
	}
}

func TestGuarded(t *testing.T) {
	t.Run("allowed runs fn and returns its result", func(t *testing.T) {
		g := NewGuard(ModeFull)
		called := false
		got, err := Guarded(g, Destructive, func() (string, error) {
			called = true
			return "ok", nil
		})
		if !called {
			t.Error("fn was not called")
		}
		if err != nil || got != "ok" {
			t.Errorf("Guarded() = (%q, %v), want (%q, nil)", got, err, "ok")
		}
	})

	t.Run("disallowed skips fn and returns ErrDisabled", func(t *testing.T) {
		g := NewGuard(ModeReadOnly)
		called := false
		got, err := Guarded(g, Destructive, func() (string, error) {
			called = true
			return "ok", nil
		})
		if called {
			t.Error("fn was called despite disallowed level")
		}
		if !errors.Is(err, ErrDisabled) {
			t.Errorf("err = %v, want ErrDisabled", err)
		}
		if got != "" {
			t.Errorf("got = %q, want zero value", got)
		}
	})
}

func levelName(l Level) string {
	switch l {
	case ReadOnly:
		return "read-only"
	case Write:
		return "write"
	case Destructive:
		return "destructive"
	default:
		return "unknown"
	}
}
