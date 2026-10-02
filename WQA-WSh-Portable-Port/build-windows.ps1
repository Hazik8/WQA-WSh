$ErrorActionPreference = 'Stop'
$Root = Split-Path -Parent $MyInvocation.MyCommand.Path
$Work = Join-Path $Root '.work'
$Src = Join-Path $Work 'WQA-WSh'
$Out = Join-Path $Root 'dist\windows-amd64'

if (-not (Get-Command git -ErrorAction SilentlyContinue)) { throw 'Git is required.' }
if (-not (Get-Command go -ErrorAction SilentlyContinue)) { throw 'Go is required.' }

New-Item -ItemType Directory -Force -Path $Work, $Out | Out-Null
if (Test-Path $Src) { Remove-Item -Recurse -Force $Src }
git clone --depth 1 https://github.com/Hazik8/WQA-WSh.git $Src

Push-Location $Src
try {
    go test ./...
    go build -trimpath -o (Join-Path $Out 'wqa.exe') ./cmd/wqa
    go build -trimpath -o (Join-Path $Out 'wsh.exe') ./cmd/wsh
    Copy-Item repository.json $Out -ErrorAction SilentlyContinue
} finally {
    Pop-Location
}

Write-Host "[OK] Windows build complete: $Out"
