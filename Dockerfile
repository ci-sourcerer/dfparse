FROM golang:1-alpine AS builder

WORKDIR /src
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /dist/dfparse ./cmd/dfparse

FROM scratch
COPY --from=builder /dist/dfparse /dfparse
ENTRYPOINT ["/dfparse"]