# `dfparse`

## Overview

`dfparse` uses [BuildKit's Dockerfile parser](https://github.com/moby/buildkit/frontend/dockerfile/parser) to dump the AST of a Dockerfile as JSON. Possible uses include calling it from another program without first-class support for the BuildKit parser (Python, anyone?), or as a debugging tool to understand how the parser is interpreting a Dockerfile.

The program reads from a file passed as the first argument.

