# `dfparse`

## Overview

`dfparse` uses [BuildKit's Dockerfile parser](https://github.com/moby/buildkit/tree/main/frontend/dockerfile/parser) to dump the AST of a Dockerfile as JSON. Possible uses include calling it from another program without first-class support for the BuildKit parser (Python, anyone?), or as a debugging tool to understand how the parser is interpreting a Dockerfile.

The program reads from a file passed as the first argument.

All images are published and tagged by release at `docker.io/cisourcerer/dfparse:<tag>` (replace `<tag>` with the desired release tag when pulling or referencing the image).
