FROM golang:1.20-bullseye as builder
ARG VCS_REF
ARG BUILD_DATE
ARG VERSION
ENV GO111MODULE="on"        \
    CGO_ENABLED=1           \
    GOOS=linux              \
    GOARCH=amd64

WORKDIR /workspace
COPY . .

# Build
RUN  go build -mod=vendor -a -o /workspace/bin/url-shortener /workspace/cmd/url-shortener/main.go
RUN chown 1001:1001 /workspace/bin/url-shortener

# Use distroless as minimal base image to package the cimgt binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/base:nonroot
WORKDIR /
COPY --from=builder /workspace/bin/url-shortener .
USER 1001:1001

ENTRYPOINT ["/url-shortener"]
