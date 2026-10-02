# WQA & WSh Portable Build Kit

This kit builds the current `Hazik8/WQA-WSh` project for Linux and Windows without changing the Windows code path.

## What it provides

- `wqa` compiler/runtime CLI
- `wsh` shell
- WQA package build/run/install flow
- repository/search/install/update functionality from the upstream project
- Windows native build (`wqa.exe`, `wsh.exe`)
- Linux native build (`wqa`, `wsh`)
- Linux user data under `~/.wsh` and `~/.windroid`
- Windows data remains compatible with the existing Windows layout

The source is fetched from the official project repository at build time so the kit does not pretend to contain a stale copy of the whole source tree.

## Linux

Requirements: Git and Go 1.26+.

```bash
chmod +x build-linux.sh
./build-linux.sh
```

Outputs are placed in `dist/linux-amd64/`.

## Windows

Requirements: Git and Go 1.26+.

Run PowerShell:

```powershell
Set-ExecutionPolicy -Scope Process Bypass
.\build-windows.ps1
```

Outputs are placed in `dist/windows-amd64/`.

## Notes

The Linux build keeps WQA/WSh functionality portable. Windows-specific commands such as `tasklist` are mapped to Linux `ps` by the patcher, and `wqa.exe` is resolved as `wqa` on Linux. The Windows build is left native and keeps `.exe` behavior.
