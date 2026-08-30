package app

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

func Run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("dfparse", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {
		fmt.Fprintln(stderr, "Usage: dfparse [options] Dockerfile")
		fmt.Fprintln(stderr, "")
		fmt.Fprintln(stderr, "Options:")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}

	if fs.NArg() != 1 {
		fmt.Fprintln(stderr, "error: Dockerfile path required")
		fs.Usage()
		return 2
	}

	f, err := os.Open(fs.Arg(0))
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	defer f.Close()

	res, err := parser.Parse(f)
	if err != nil {
		fmt.Fprintln(stderr, "parse:", err)
		return 1
	}

	if err := json.NewEncoder(stdout).Encode(res.AST); err != nil {
		fmt.Fprintln(stderr, "encode:", err)
		return 1
	}
	return 0
}
