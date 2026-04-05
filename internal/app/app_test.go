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

	return f.Name()
}

func TestRun_DumpsAST(t *testing.T) {
	path := writeTempDockerfile(t, "FROM alpine\n")
	defer os.Remove(path)

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
