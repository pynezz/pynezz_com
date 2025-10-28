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
