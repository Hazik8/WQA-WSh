#!/usr/bin/env python3
import pathlib
import sys

if len(sys.argv) != 2:
    raise SystemExit('usage: patch-linux.py <WQA-WSh-source>')

root = pathlib.Path(sys.argv[1])

# 1. Make WSh use the normal home directory on Linux for its cache files.
for rel in ['cmd/wsh/aliases.go', 'cmd/wsh/history.go']:
    p = root / rel
    s = p.read_text(encoding='utf-8')
    s = s.replace('userProfile := os.Getenv("USERPROFILE")\n\n\tif userProfile == "" {', 'userProfile, _ := os.UserHomeDir()\n\n\tif userProfile == "" {')
    p.write_text(s, encoding='utf-8')

# 2. Portable repository cache paths in commands.go.
p = root / 'cmd/wsh/commands.go'
s = p.read_text(encoding='utf-8')
s = s.replace('userProfile := os.Getenv("USERPROFILE")\n\tif userProfile == "" {\n\t\treturn fmt.Errorf(\n\t\t\t"USERPROFILE environment variable is not set",', 'userProfile, homeErr := os.UserHomeDir()\n\tif homeErr != nil || userProfile == "" {\n\t\treturn fmt.Errorf(\n\t\t\t"cannot determine user home directory: %v", homeErr,')
s = s.replace('userProfile := os.Getenv("USERPROFILE")\n\tif userProfile == "" {', 'userProfile, _ := os.UserHomeDir()\n\tif userProfile == "" {')
s = s.replace('cmd := exec.Command("tasklist")', 'cmdName := "ps"\n\tcmdArgs := []string{"-eo", "pid,comm"}\n\tif os.PathSeparator == \'\\\\\' {\n\t\tcmdName = "tasklist"\n\t\tcmdArgs = nil\n\t}\n\tcmd := exec.Command(cmdName, cmdArgs...)')
s = s.replace('target := "wqa.exe"', 'target := "wqa"\n\tif os.PathSeparator == \'\\\\\' {\n\t\ttarget = "wqa.exe"\n\t}')
p.write_text(s, encoding='utf-8')

# 3. Make the package installer choose the native platform path.
p = root / 'internal/installer/installer.go'
s = p.read_text(encoding='utf-8')
s = s.replace('"path/filepath"', '"path/filepath"\n\t"runtime"')
s = s.replace('appsPath := paths.GetAppsPath(\n\t\tpaths.Windows,\n\t)', 'platform := paths.Platform("linux")\n\tif runtime.GOOS == "windows" {\n\t\tplatform = paths.Windows\n\t}\n\tappsPath := paths.GetAppsPath(platform)')
p.write_text(s, encoding='utf-8')

# 4. Add a proper Linux application path and keep Windows unchanged.
p = root / 'internal/paths/paths.go'
s = p.read_text(encoding='utf-8')
s = s.replace('const (\n\tWindows Platform = "windows"\n)', 'const (\n\tWindows Platform = "windows"\n\tLinux Platform = "linux"\n)')
s = s.replace('case Windows:\n\t\treturn `C:\\\\WinDroid\\\\Apps`\n\n\tdefault:', 'case Windows:\n\t\treturn `C:\\\\WinDroid\\\\Apps`\n\tcase Linux:\n\t\thome, err := os.UserHomeDir()\n\t\tif err != nil || home == "" {\n\t\t\treturn filepath.Join(".", "WinDroid", "Apps")\n\t\t}\n\t\treturn filepath.Join(home, ".windroid", "Apps")\n\n\tdefault:')
s = s.replace('projectPath := filepath.Join(\n\t\t`C:\\\\WQA`,\n\t\t"repository.json",\n\t)', 'projectPath := filepath.Join(`C:\\\\WQA`, "repository.json")')
p.write_text(s, encoding='utf-8')

print('[OK] Linux portability patches applied')
