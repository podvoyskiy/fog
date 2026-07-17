package cmds

import (
	"fmt"
	"testing"

	"github.com/podvoyskiy/fog/config"
)

func TestHandleCmd(t *testing.T) {
	tmpDir := t.TempDir()
	config, err := config.Load(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		arg     []string
		wantErr bool
	}{
		{[]string{"fake arg"}, true},
		{[]string{"-h"}, false},
		{[]string{"-l", "foo"}, true},
		{[]string{"--limit", "30"}, false},
		{[]string{"--freq-bonus", "256"}, true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			err := HandleCmd(config, tt.arg)
			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestHandleCmdUpdatesConfig(t *testing.T) {
	tmpDir := t.TempDir()
	cfg, err := config.Load(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	if err = HandleCmd(cfg, []string{"-l", "20"}); err != nil {
		t.Fatal(err)
	}
	if err = HandleCmd(cfg, []string{"-b", "0"}); err != nil {
		t.Fatal(err)
	}

	newCfg, err := config.Load(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	if newCfg.Limit != 20 {
		t.Fatalf("config not saved: limit = %d, want 20", newCfg.Limit)
	}
	if newCfg.MaxFreqBonus != 0 {
		t.Fatalf("config not saved: max_freq_bonus = %d, want 0", newCfg.MaxFreqBonus)
	}
}
