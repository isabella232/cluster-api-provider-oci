# Build the manager binary
FROM golang:1.17.3 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

ARG package=.
ARG ARCH
ARG LDFLAGS

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY cloud/ cloud/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -ldflags "${LDFLAGS} -extldflags '-static'"  -o manager ${package}

FROM ghcr.io/oracle/oraclelinux:8-slim
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]