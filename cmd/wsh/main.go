package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"windroid/wqa/internal/installer"
	"windroid/wqa/internal/packages"
	"windroid/wqa/internal/repository"
)

const version = "0.4.0"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("WinDroid Shell", version)
	fmt.Println("------------------------------")
	fmt.Println("Type 'help' for available commands.")
	fmt.Println()

	for {
		dir, err := os.Getwd()
		if err != nil {
			dir = "?"
		}

		fmt.Printf("wsh [%s]> ", dir)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println()
			return
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if !handleCommand(input) {
			return
		}
	}
}

func handleCommand(input string) bool {
	parts := strings.Fields(input)

	if len(parts) == 0 {
		return true
	}

	command := strings.ToLower(parts[0])
	args := parts[1:]

	switch command {

	case "help":
		printHelp()

	case "version":
		fmt.Println("WinDroid Shell", version)

	case "clear", "cls":
		clearScreen()

	case "exit", "quit":
		fmt.Println("Goodbye.")
		return false

	case "pwd":
		printWorkingDirectory()

	case "cd":
		changeDirectory(args)

	case "dir", "ls":
		listDirectory(args)

	case "type", "cat":
		showFile(args)

	case "mkdir":
		makeDirectory(args)

	case "del", "rm":
		deleteFile(args)

	case "run":
		runApplication(args)

	case "wqa":
		runWQA(args)

	case "install":
		if len(args) == 0 {
			fmt.Println("Usage: install <package>")
			break
		}

		if err := installApplication(args[0]); err != nil {
			fmt.Println("[ERROR]", err)
		}

	case "remove":
		if len(args) == 0 {
			fmt.Println("Usage: remove <app>")
			break
		}

		if err := packages.Remove(args[0]); err != nil {
			fmt.Println("[ERROR]", err)
		}

	case "list":
		if err := packages.List(); err != nil {
			fmt.Println("[ERROR]", err)
		}

	case "info":
		if len(args) == 0 {
			fmt.Println("Usage: info <app>")
			break
		}

		if err := packages.Info(args[0]); err != nil {
			fmt.Println("[ERROR]", err)
		}

	case "search":
		searchRepository(args)

	default:
		runExternal(command, args)
	}

	return true
}

func printHelp() {
	fmt.Println(`
WinDroid Shell

File commands:

  pwd
      Show current directory

  cd <path>
      Change directory

  dir
      List files

  type <file>
      Show file contents

  mkdir <name>
      Create directory

  del <file>
      Delete file

  info <app>
      Show application information

  search [name]
      Search installed applications

WQA:

  wqa <command>
      Run WQA CLI

      Examples:
        wqa list
        wqa install Calculator.wqa
        wqa run Calculator.wqa
        wqa info Calculator.wqa

Shell:

  help
      Show this help

  version
      Show WSh version

  clear
      Clear screen

  exit
      Exit WSh

Package commands:

  install <file>
      Install a supported package

  remove <app>
      Remove an installed application

  list
      List installed applications

  run <program.exe>
      Run a Windows EXE
 
 `)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printWorkingDirectory() {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(dir)
}

func changeDirectory(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: cd <path>")
		return
	}

	path := strings.Join(args, " ")

	if err := os.Chdir(path); err != nil {
		fmt.Println("Error:", err)
	}
}

func listDirectory(args []string) {
	path := "."

	if len(args) > 0 {
		path = strings.Join(args, " ")
	}

	entries, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("[DIR]  %s\n", entry.Name())
		} else {
			fmt.Printf("       %s\n", entry.Name())
		}
	}
}

func showFile(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: type <file>")
		return
	}

	path := strings.Join(args, " ")

	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print(string(data))
}

func makeDirectory(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: mkdir <name>")
		return
	}

	name := strings.Join(args, " ")

	if err := os.MkdirAll(name, 0755); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("[OK] Directory created:", name)
}

func deleteFile(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: del <file>")
		return
	}

	path := strings.Join(args, " ")

	if err := os.Remove(path); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("[OK] Deleted:", path)
}

func runWQA(args []string) {
	wqa, err := findWQA()

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	cmd := exec.Command(wqa, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println("[ERROR]", err)
	}
}

func findWQA() (string, error) {
	// 1. Ищем wqa.exe рядом с wsh.exe.
	exe, err := os.Executable()

	if err == nil {
		dir := filepath.Dir(exe)

		local := filepath.Join(dir, "wqa.exe")

		if _, err := os.Stat(local); err == nil {
			return local, nil
		}
	}

	// 2. Ищем через PATH.
	path, err := exec.LookPath("wqa.exe")

	if err == nil {
		return path, nil
	}

	return "", fmt.Errorf(
		"wqa.exe not found; build WQA or add it to PATH",
	)
}

func runExternal(command string, args []string) {
	exe, err := exec.LookPath(command)

	if err == nil {
		cmd := exec.Command(exe, args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}

		return
	}

	// Передаём неизвестную команду Windows CMD.
	all := append([]string{command}, args...)

	cmd := exec.Command(
		"cmd",
		append([]string{"/c"}, all...)...,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println("Command not found:", command)
	}
}

func runApplication(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: run <app|file>")
		return
	}

	target := strings.Join(args, " ")

	// --------------------------------
	// 1. WQA
	// --------------------------------

	if strings.HasSuffix(
		strings.ToLower(target),
		".wqa",
	) {
		fmt.Println("[INFO] WQA application")

		runWQA([]string{
			"run",
			target,
		})

		return
	}

	// --------------------------------
	// 2. Точный путь
	// --------------------------------

	if strings.Contains(target, `\`) ||
		strings.Contains(target, `/`) {

		if _, err := os.Stat(target); err == nil {
			runEXEFile(target)
			return
		}
	}

	// --------------------------------
	// 3. EXE в текущей папке
	// --------------------------------

	localPath := filepath.Join(".", target)

	if info, err := os.Stat(localPath); err == nil {
		if !info.IsDir() {
			absolutePath, err := filepath.Abs(localPath)

			if err != nil {
				fmt.Println("[ERROR]", err)
				return
			}

			runEXEFile(absolutePath)
			return
		}
	}

	// --------------------------------
	// 4. Установленное приложение
	// --------------------------------

	appName := strings.TrimSuffix(
		target,
		filepath.Ext(target),
	)

	appPath := filepath.Join(
		packages.AppsDir,
		appName,
		"app.exe",
	)

	if _, err := os.Stat(appPath); err == nil {
		fmt.Println("[INFO] Installed application")
		runEXEFile(appPath)
		return
	}

	// --------------------------------
	// 5. Windows PATH
	// --------------------------------

	if resolved, err := exec.LookPath(target); err == nil {
		runEXEFile(resolved)
		return
	}

	fmt.Println("[ERROR] Application not found:", target)
}

func runEXEFile(target string) {
	// Если это просто имя программы,
	// ищем её через Windows PATH.
	resolved, err := exec.LookPath(target)

	if err == nil {
		target = resolved
	} else {
		// Если это путь к конкретному EXE,
		// проверяем его напрямую.
		if _, statErr := os.Stat(target); statErr != nil {
			fmt.Println("[ERROR]", statErr)
			return
		}
	}

	cmd := exec.Command(target)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("[INFO] Starting:", target)

	if err := cmd.Start(); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] Process started")
}

func searchRepository(args []string) {
	query := strings.Join(args, " ")

	repo, err := repository.Load("repository.json")

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	results := repo.Search(query)

	fmt.Println("WinDroid Repository")
	fmt.Println("-------------------")

	if len(results) == 0 {
		fmt.Println("No applications found")
		return
	}

	for _, app := range results {
		fmt.Printf(
			"%-24s %-12s %s\n",
			app.Name,
			app.Type,
			app.Version,
		)
	}
}

func showAppInfo(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: info <app>")
		return
	}

	query := strings.Join(args, " ")

	// Сначала ищем среди установленных приложений.
	if err := packages.Info(query); err == nil {
		return
	}

	// Если приложение не установлено,
	// ищем его в репозитории.
	repo, err := repository.Load("repository.json")

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	app := repo.Find(query)

	if app == nil {
		fmt.Println("[ERROR] Application not found:", query)
		return
	}

	fmt.Println("Application Information")
	fmt.Println("-----------------------")
	fmt.Println("Name:", app.Name)
	fmt.Println("ID:", app.ID)
	fmt.Println("Version:", app.Version)
	fmt.Println("Type:", app.Type)

	if app.Download != "" {
		fmt.Println("Download:", app.Download)
	} else {
		fmt.Println("Download: not available")
	}
}

func installApplication(target string) error {
	ext := strings.ToLower(filepath.Ext(target))

	// Локальная установка WQA или EXE.
	if ext == ".wqa" || ext == ".exe" {
		return packages.Install(target)
	}

	// Загружаем каталог репозитория.
	repo, err := repository.Load("repository.json")
	if err != nil {
		return fmt.Errorf(
			"cannot load repository: %w",
			err,
		)
	}

	app := repo.Find(target)

	if app == nil {
		return fmt.Errorf(
			"application not found: %s",
			target,
		)
	}

	fmt.Println("[INFO] Found:", app.Name)
	fmt.Println("[INFO] Version:", app.Version)
	fmt.Println("[INFO] Type:", app.Type)

	if app.Download == "" {
		return fmt.Errorf(
			"application has no download URL",
		)
	}

	// Пока онлайн-установка поддерживает WQA.
	if strings.ToLower(app.Type) != "wqa" {
		return fmt.Errorf(
			"unsupported repository package type: %s",
			app.Type,
		)
	}

	tempDir, err := os.MkdirTemp("", "wqa-install-*")
	if err != nil {
		return err
	}

	defer os.RemoveAll(tempDir)

	packagePath := filepath.Join(
		tempDir,
		app.Name+".wqa",
	)

	fmt.Println("[INFO] Downloading...")

	if err := repository.Download(
		app,
		packagePath,
	); err != nil {
		return err
	}

	fmt.Println("[INFO] Verifying SHA-256...")

	if err := repository.VerifySHA256(
		packagePath,
		app.SHA256,
	); err != nil {
		return err
	}

	fmt.Println("[OK] SHA-256 verified")

	fmt.Println("[INFO] Installing...")

	if err := installer.Install(packagePath); err != nil {
		return fmt.Errorf(
			"installation failed: %w",
			err,
		)
	}

	return nil
}
