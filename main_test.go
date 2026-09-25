package main

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantExit   int
		wantStdout string
		wantStderr string
	}{
		{
			name:       "empty args prints usage with exit 2",
			args:       []string{},
			wantExit:   ExitUsage,
			wantStdout: "Novel Translator CLI",
		},
		{
			name:       "help command exits 0",
			args:       []string{"help"},
			wantExit:   ExitSuccess,
			wantStdout: "Usage:",
		},
		{
			name:       "-help flag exits 0",
			args:       []string{"-help"},
			wantExit:   ExitSuccess,
			wantStdout: "Usage:",
		},
		{
			name:       "unknown command exits 2",
			args:       []string{"unknown-subcommand"},
			wantExit:   ExitUsage,
			wantStderr: "unknown command",
		},
		{
			name:       "invalid flag exits 2",
			args:       []string{"validate", "-invalid-flag-123"},
			wantExit:   ExitUsage,
			wantStderr: "flag provided but not defined",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			ctx := context.Background()

			exitCode := run(ctx, tt.args, &stdout, &stderr)
			if exitCode != tt.wantExit {
				t.Errorf("run(%v) exitCode = %d, want %d", tt.args, exitCode, tt.wantExit)
			}

			if tt.wantStdout != "" && !strings.Contains(stdout.String(), tt.wantStdout) {
				t.Errorf("stdout = %q, expected to contain %q", stdout.String(), tt.wantStdout)
			}
			if tt.wantStderr != "" && !strings.Contains(stderr.String(), tt.wantStderr) {
				t.Errorf("stderr = %q, expected to contain %q", stderr.String(), tt.wantStderr)
			}
		})
	}
}
