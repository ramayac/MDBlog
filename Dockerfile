# ── Stage 1: Build Go binaries ───────────────────────────────────────────────
FROM golang:1.24 AS build

WORKDIR /src

# Download dependencies first (cached layer)
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION=dev
ARG COMMIT=unknown
ARG DATE=unknown
ARG LDFLAGS="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}"

# Build the Lambda entry-point binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "${LDFLAGS}" -o /out/lambda ./cmd/lambda

# Build the CLI tool so we can run build-index
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /out/mdblog ./cmd/mdblog

# Generate the post metadata index (baked into the image)
RUN cd /src && /out/mdblog build-index

# ── Stage 2: Minimal production image ────────────────────────────────────────
FROM scratch

# CA certificates (needed for any outbound TLS from the binary, e.g. latency checks)
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /out/lambda  /lambda
COPY --from=build /out/mdblog  /mdblog
COPY --from=build /src/posts/      /posts/
COPY --from=build /src/assets/     /assets/
COPY --from=build /src/templates/  /templates/
COPY --from=build /src/config.toml /config.toml

ENTRYPOINT ["/lambda"]