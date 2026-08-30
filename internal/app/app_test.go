package app

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func writeTempDockerfile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "dockerfile-*.Dockerfile")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("write temp Dockerfile: %v", err)
	}

	t.Cleanup(func() { os.Remove(f.Name()) })
	return f.Name()
}

func TestRun_DumpsAST(t *testing.T) {
	path := writeTempDockerfile(t, "FROM alpine\n")

	var stdout, stderr bytes.Buffer
	code := Run([]string{path}, &stdout, &stderr)

	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
	if stderr.Len() != 0 {
		t.Fatalf("expected no stderr output, got %q", stderr.String())
	}
	if !json.Valid(bytes.TrimSpace(stdout.Bytes())) {
		t.Fatalf("expected valid JSON output, got %q", stdout.String())
	}
}

func TestRun_Help(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run([]string{"--help"}, &stdout, &stderr)

	if code != 0 {
		t.Fatalf("expected exit 0 for help, got %d", code)
	}
	if stdout.Len() != 0 {
		t.Fatalf("expected no stdout output for help, got %q", stdout.String())
	}
	if !strings.Contains(stderr.String(), "Usage: dfparse") {
		t.Fatalf("expected usage output on stderr, got %q", stderr.String())
	}
}

func TestRun_Errors(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantCode int
	}{
		{"no args", []string{}, 2},
		{"too many args", []string{"a", "b"}, 2},
		{"unrecognised flag", []string{"--foo"}, 2},
		{"nonexistent file", []string{"/nonexistent/path/Dockerfile"}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := Run(tt.args, &stdout, &stderr)
			if code != tt.wantCode {
				t.Errorf("expected exit %d, got %d (stderr: %q)", tt.wantCode, code, stderr.String())
			}
			if stdout.Len() != 0 {
				t.Errorf("expected no stdout output, got %q", stdout.String())
			}
		})
	}
}
