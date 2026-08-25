package setup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const installDir = `C:\WinDroid\Tools`

func Setup() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	exe, err = filepath.Abs(exe)
	if err != nil {
		return err
	}

	fmt.Println("[INFO] WQA Windows Setup")
	fmt.Println("[INFO] Installing to:", installDir)

	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("cannot create install directory: %w", err)
	}

	target := filepath.Join(installDir, "wqa.exe")

	if strings.EqualFold(exe, target) {
		fmt.Println("[INFO] WQA is already installed")
	} else {
		data, err := os.ReadFile(exe)
		if err != nil {
			return err
		}

		if err := os.WriteFile(target, data, 0755); err != nil {
			return fmt.Errorf("cannot copy wqa.exe: %w", err)
		}

		fmt.Println("[OK] Installed:", target)
	}

	if err := addToPath(installDir); err != nil {
		return err
	}

	fmt.Println("[OK] PATH configured")
	fmt.Println()
	fmt.Println("WQA setup completed.")
	fmt.Println("Open a new terminal and run:")
	fmt.Println()
	fmt.Println("  wqa --help")

	return nil
}

func addToPath(dir string) error {
	cmd := exec.Command(
		"powershell",
		"-NoProfile",
		"-Command",
		`$old = [Environment]::GetEnvironmentVariable("Path", "User"); `+
			`if ([string]::IsNullOrWhiteSpace($old)) { $old = "" }; `+
			`$items = $old -split ';' | Where-Object { $_ -ne "" }; `+
			`if ($items -notcontains $env:WQA_INSTALL_DIR) { `+
			`[Environment]::SetEnvironmentVariable("Path", (($items + $env:WQA_INSTALL_DIR) -join ';'), "User") `+
			`}`,
	)

	cmd.Env = append(
		os.Environ(),
		"WQA_INSTALL_DIR="+dir,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"cannot update PATH: %s: %w",
			strings.TrimSpace(string(output)),
			err,
		)
	}

	return nil
}
