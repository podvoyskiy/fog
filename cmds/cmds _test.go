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
		{[]string{"-f", "10"}, true},
		{[]string{"-f", "3"}, false},
		{[]string{"--max_results", "30"}, false},
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

	err = HandleCmd(cfg, []string{"-m", "20"})
	if err != nil {
		t.Fatal(err)
	}

	newCfg, err := config.Load(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	if newCfg.MaxResults != 20 {
		t.Fatalf("config not saved: MaxResults = %d, want 20", newCfg.MaxResults)
	}
}
