FROM golang:1.25-alpine3.21 AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /dist/dfparse ./cmd/dfparse

FROM scratch
COPY --from=builder /dist/dfparse /dfparse
USER 65534:65534
ENTRYPOINT ["/dfparse"]