#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")" && pwd)"
WORK="$ROOT/.work"
SRC="$WORK/WQA-WSh"
OUT="$ROOT/dist/linux-amd64"

command -v git >/dev/null || { echo "Git is required."; exit 1; }
command -v go >/dev/null || { echo "Go is required."; exit 1; }

mkdir -p "$WORK" "$OUT"
rm -rf "$SRC"
git clone --depth 1 https://github.com/Hazik8/WQA-WSh.git "$SRC"

python3 "$ROOT/patch-linux.py" "$SRC"

cd "$SRC"
go test ./...
go build -trimpath -o "$OUT/wqa" ./cmd/wqa
go build -trimpath -o "$OUT/wsh" ./cmd/wsh

cp README.md "$OUT/UPSTREAM-README.md" 2>/dev/null || true
cp repository.json "$OUT/repository.json" 2>/dev/null || true
chmod +x "$OUT/wqa" "$OUT/wsh"

echo
printf '[OK] Linux build complete: %s\n' "$OUT"
