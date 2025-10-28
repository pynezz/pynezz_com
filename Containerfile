FROM registry.access.redhat.com/ubi10/nodejs-22@sha256:f8a43e16bc579770605ba1512522a5f70586dd8a8f13924d48e0f0ac61cabb17 AS assets
WORKDIR /workspace

USER root

COPY --chown=1000:1000 package.json package-lock.json ./
RUN chown -R 1000:1000 "/workspace"
COPY --chown=1000:1000 . .

USER 1000

RUN npm ci && ( npm audit fix || true )
RUN npx update-browserslist-db@latest --yes || true
RUN npm run build:css

USER root


## ---------------------  Build step ----------------------------------------- ##
FROM --platform=amd64 registry.access.redhat.com/ubi10/go-toolset:1.24 AS builder
WORKDIR /app
USER root
ENV GOTOOLCHAIN=auto
ENV GOPATH=/go
ENV PATH=$PATH:/opt/app-root/src/.local/bin
ARG ZIG_VERSION=0.12.0
ARG ZIG_ARCHIVE=zig-linux-x86_64-${ZIG_VERSION}.tar.xz

RUN dnf update -y && \
    dnf install -y --nodocs \
      git \
      make \
      ca-certificates \
      xz \
      curl && \
    curl -L "https://ziglang.org/download/${ZIG_VERSION}/${ZIG_ARCHIVE}" -o /tmp/${ZIG_ARCHIVE} && \
    mkdir -p /opt/zig && \
    tar -xJf /tmp/${ZIG_ARCHIVE} -C /opt/zig --strip-components=1 && \
    ln -sf /opt/zig/zig /usr/local/bin/zig && \
    rm -rf /tmp/${ZIG_ARCHIVE} /var/cache/dnf/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=assets /workspace/pynezz/public/css/styles.css ./pynezz/public/css/styles.css

RUN go install github.com/a-h/templ/cmd/templ@v0.3.960 && \
    /go/bin/templ generate

ARG BIN_NAME=pynezz-cli
RUN VERSION=$(git describe --tags --always --long 2>/dev/null || echo dev) && \
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
    CC="zig cc -target x86_64-linux-gnu.2.31.0" \
    CXX="zig c++ -target x86_64-linux-gnu.2.31.0" \
    go build -o "${BIN_NAME}_linux_amd64.out" \
      -tags linux \
      -ldflags="-s -w -X main.buildVersion=${VERSION}" .

FROM --platform=amd64 registry.access.redhat.com/ubi10-minimal:latest AS runtime
WORKDIR /app

RUN microdnf update -y && \
    microdnf install -y ca-certificates && \
    microdnf clean all

COPY --from=builder /app/pynezz-cli_linux_amd64.out ./pynezz-cli_linux_amd64.out
COPY --from=builder /app/pynezz ./pynezz
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/content ./content
COPY --from=builder /app/config ./config

RUN cat <<'EOF' > /app/entrypoint.sh
#!/usr/bin/env sh
set -eu

auto_parse() {
  echo "[pynezz] parsing markdown content..."
  /app/pynezz-cli_linux_amd64.out cms parse
}

case "${PYNEZZ_AUTO_PARSE:-0}" in
  1|true|TRUE|on|ON)
    auto_parse
    ;;
esac

exec /app/pynezz-cli_linux_amd64.out "$@"
EOF
RUN chmod +x /app/entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["serve"]
