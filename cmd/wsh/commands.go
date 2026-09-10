package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"windroid/wqa/internal/installer"
	"windroid/wqa/internal/packages"
	"windroid/wqa/internal/repository"
	"windroid/wqa/internal/updater"
)

func parseCommandLine(input string) []string {
	var result []string
	var current strings.Builder

	var quote rune
	escaped := false

	for _, r := range input {
		if escaped {
			current.WriteRune(r)
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		if quote != 0 {
			if r == quote {
				quote = 0
			} else {
				current.WriteRune(r)
			}
			continue
		}

		switch r {
		case '\'', '"':
			quote = r

		case ' ', '\t':
			if current.Len() > 0 {
				result = append(result, current.String())
				current.Reset()
			}

		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printWorkingDirectory() {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("[ERROR]", err)
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
		fmt.Println("[ERROR]", err)
	}
}

func listDirectory(args []string) {
	path := "."

	if len(args) > 0 {
		path = strings.Join(args, " ")
	}

	entries, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("[DIR]  %s\n", entry.Name())
			continue
		}

		info, err := entry.Info()

		if err != nil {
			fmt.Printf("       %s\n", entry.Name())
			continue
		}

		fmt.Printf(
			"       %-30s %d bytes\n",
			entry.Name(),
			info.Size(),
		)
	}
}

func showFile(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: cat <file>")
		return
	}

	path := strings.Join(args, " ")

	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("[ERROR]", err)
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
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] Directory created:", name)
}

func newItem(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: new <file>")
		return
	}

	path := strings.Join(args, " ")

	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_EXCL,
		0644,
	)

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	file.Close()

	fmt.Println("[OK] File created:", path)
}

func deleteFile(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: del <file>")
		return
	}

	path := strings.Join(args, " ")

	if err := os.Remove(path); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] Deleted:", path)
}

func copyItem(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: cp <source> <destination>")
		return
	}

	source := args[0]
	destination := args[1]

	data, err := os.ReadFile(source)

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	if err := os.WriteFile(destination, data, 0644); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] Copied:", source, "->", destination)
}

func moveItem(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: mv <source> <destination>")
		return
	}

	if err := os.Rename(args[0], args[1]); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] Moved:", args[0], "->", args[1])
}

func writeFile(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: write <file> <text>")
		return
	}

	path := args[0]
	text := strings.Join(args[1:], " ")

	if err := os.WriteFile(
		path,
		[]byte(text),
		0644,
	); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] File written:", path)
}

func appendFile(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: append <file> <text>")
		return
	}

	path := args[0]
	text := strings.Join(args[1:], " ")

	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	defer file.Close()

	if _, err := file.WriteString(text + "\n"); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] Text appended:", path)
}

func testPath(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: test <path>")
		return
	}

	path := strings.Join(args, " ")

	info, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("False")
			return
		}

		fmt.Println("[ERROR]", err)
		return
	}

	if info.IsDir() {
		fmt.Println("True (directory)")
	} else {
		fmt.Println("True (file)")
	}
}

func listProcesses() {
	cmd := exec.Command("tasklist")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("[ERROR]", err)
	}
}

func killProcess(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: kill <pid>")
		return
	}

	pid, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("[ERROR] Invalid PID:", args[0])
		return
	}

	process, err := os.FindProcess(pid)

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	if err := process.Kill(); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] Process terminated:", pid)
}

func showEnvironment(args []string) {
	if len(args) == 0 {
		for _, value := range os.Environ() {
			fmt.Println(value)
		}

		return
	}

	name := args[0]

	value, exists := os.LookupEnv(name)

	if !exists {
		fmt.Println("[INFO] Environment variable not found:", name)
		return
	}

	fmt.Println(value)
}

func setVariable(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: set <name> <value>")
		return
	}

	name := args[0]
	value := strings.Join(args[1:], " ")

	if err := os.Setenv(name, value); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Printf("[OK] %s=%s\n", name, value)
}

func unsetVariable(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: unset <name>")
		return
	}

	name := args[0]

	if err := os.Unsetenv(name); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] Removed:", name)
}

func showDate() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func whichCommand(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: which <command>")
		return
	}

	path, err := exec.LookPath(args[0])

	if err != nil {
		fmt.Println("[INFO] Command not found:", args[0])
		return
	}

	fmt.Println(path)
}

func echoCommand(args []string) {
	if len(args) == 0 {
		fmt.Println()
		return
	}

	text := strings.Join(args, " ")

	if text == "$?" {
		fmt.Println(lastExitCode)
		return
	}

	fmt.Println(text)
}

func runExternal(command string, args []string) {
	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			lastExitCode = exitErr.ExitCode()
		} else {
			lastExitCode = 1
		}

		fmt.Println("[ERROR]", err)
		return
	}

	lastExitCode = 0
}

func runApplication(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: run <program.exe>")
		return
	}

	target := args[0]
	extraArgs := args[1:]

	// Direct path.
	if _, err := os.Stat(target); err == nil {
		startProcess(target, extraArgs)
		return
	}

	// PATH lookup.
	path, err := exec.LookPath(target)

	if err == nil {
		startProcess(path, extraArgs)
		return
	}

	// Windows system directory.
	if !filepath.IsAbs(target) {
		systemRoot := os.Getenv("SystemRoot")

		if systemRoot != "" {
			systemPath := filepath.Join(
				systemRoot,
				"System32",
				target,
			)

			if _, err := os.Stat(systemPath); err == nil {
				startProcess(systemPath, extraArgs)
				return
			}
		}
	}

	fmt.Println("[ERROR] Application not found:", target)
}

func startProcess(path string, args []string) {
	cmd := exec.Command(path, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("[ERROR]", err)
		lastExitCode = 1
		return
	}

	fmt.Println("[OK] Process started:", path)
	lastExitCode = 0
}

func runWQA(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: wqa <command>")
		return
	}

	executable := "wqa"

	cmd := exec.Command(executable, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		lastExitCode = 1
		fmt.Println("[ERROR]", err)
		return
	}

	lastExitCode = 0
}

func handleRepo(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: repo update")
		lastExitCode = 0
		return
	}

	switch strings.ToLower(args[0]) {
	case "update":
		if err := updateRepository(); err != nil {
			fmt.Println("[ERROR]", err)
			lastExitCode = 1
			return
		}

		lastExitCode = 0

	default:
		fmt.Println("Usage: repo update")
		lastExitCode = 1
	}
}

func handleInstall(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: install <package|app>")
		return
	}

	if err := installApplication(args[0]); err != nil {
		fmt.Println("[ERROR]", err)
		lastExitCode = 1
		return
	}

	lastExitCode = 0
}

func handleUpdate(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: update <app|all>")
		return
	}

	if strings.EqualFold(args[0], "all") {
		if err := updater.UpdateAll(); err != nil {
			fmt.Println("[ERROR]", err)
			lastExitCode = 1
			return
		}

		lastExitCode = 0
		return
	}

	if err := updater.Update(args[0]); err != nil {
		fmt.Println("[ERROR]", err)
		lastExitCode = 1
		return
	}

	lastExitCode = 0
}

func handleRemoveApp(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: remove-app <app>")
		return
	}

	if err := packages.Remove(args[0]); err != nil {
		fmt.Println("[ERROR]", err)
		lastExitCode = 1
		return
	}

	lastExitCode = 0
}

func handlePackageList() {
	if err := packages.List(); err != nil {
		fmt.Println("[ERROR]", err)
		lastExitCode = 1
		return
	}

	lastExitCode = 0
}

func updateRepository() error {
	fmt.Println("[INFO] Updating application repository...")
	fmt.Println("[INFO] Downloading repository...")

	repo, err := repository.LoadURL(
		repository.DefaultRepositoryURL,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to update repository: %w",
			err,
		)
	}

	userProfile := os.Getenv("USERPROFILE")
	if userProfile == "" {
		return fmt.Errorf(
			"USERPROFILE environment variable is not set",
		)
	}

	cachePath := filepath.Join(
		userProfile,
		".wsh",
		"repository.json",
	)

	if err := repository.Save(
		cachePath,
		repo,
	); err != nil {
		return err
	}

	fmt.Println("[OK] Repository updated.")
	fmt.Printf(
		"[INFO] Applications: %d\n",
		len(repo.Apps),
	)

	return nil
}

func loadApplicationRepository() (*repository.Repository, error) {
	// 1. Пытаемся загрузить официальный каталог с GitHub.
	repo, err := repository.LoadURL(
		repository.DefaultRepositoryURL,
	)

	if err == nil {
		// Сохраняем успешную загрузку в локальный кэш.
		if userProfile := os.Getenv("USERPROFILE"); userProfile != "" {
			cachePath := filepath.Join(
				userProfile,
				".wsh",
				"repository.json",
			)

			if err := repository.Save(
				cachePath,
				repo,
			); err != nil {
				fmt.Println(
					"[WARN] Failed to save repository cache:",
					err,
				)
			}
		}

		return repo, nil
	}

	// 2. GitHub недоступен, пробуем локальный кэш.
	fmt.Println(
		"[WARN] Failed to load online repository.",
	)

	if userProfile := os.Getenv("USERPROFILE"); userProfile != "" {
		cachePath := filepath.Join(
			userProfile,
			".wsh",
			"repository.json",
		)

		if _, statErr := os.Stat(cachePath); statErr == nil {
			fmt.Println(
				"[INFO] Using local repository cache.",
			)

			cachedRepo, cacheErr := repository.Load(
				cachePath,
			)

			if cacheErr == nil {
				return cachedRepo, nil
			}
		}
	}

	return nil, fmt.Errorf(
		"failed to load application repository: %w",
		err,
	)
}

func installApplication(target string) error {
	repo, err := loadApplicationRepository()
	if err != nil {
		return err
	}

	app := repo.Find(target)

	if app == nil {
		return fmt.Errorf(
			"application not found in repository: %s",
			target,
		)
	}

	if app.Download == "" {
		return fmt.Errorf(
			"application has no download URL",
		)
	}

	fmt.Println(
		"[INFO] Found application:",
		app.Name,
	)

	fmt.Println(
		"[INFO] Version:",
		app.Version,
	)

	fmt.Println("[INFO] Downloading...")

	tempDir, err := os.MkdirTemp(
		"",
		"wsh-install-*",
	)
	if err != nil {
		return err
	}

	defer os.RemoveAll(tempDir)

	packagePath := filepath.Join(
		tempDir,
		filepath.Base(app.Download),
	)

	if err := repository.Download(
		app,
		packagePath,
	); err != nil {
		return err
	}

	fmt.Println("[INFO] Download complete.")

	if err := repository.VerifySHA256(
		packagePath,
		app.SHA256,
	); err != nil {
		return fmt.Errorf(
			"package verification failed: %w",
			err,
		)
	}

	fmt.Println("[OK] SHA-256 verified.")

	if err := installer.Install(
		packagePath,
	); err != nil {
		return err
	}

	return nil
}

func findRepositoryFile() (string, error) {
	if userProfile := os.Getenv("USERPROFILE"); userProfile != "" {
		cachePath := filepath.Join(
			userProfile,
			".wsh",
			"repository.json",
		)

		if _, err := os.Stat(cachePath); err == nil {
			return cachePath, nil
		}
	}

	return "", fmt.Errorf(
		"repository cache not found",
	)
}

func searchRepository(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: search <query>")
		lastExitCode = 1
		return
	}

	query := strings.Join(
		args,
		" ",
	)

	repo, err := loadApplicationRepository()
	if err != nil {
		fmt.Println(
			"[ERROR]",
			err,
		)

		lastExitCode = 1
		return
	}

	results := repo.Search(query)

	if len(results) == 0 {
		fmt.Printf(
			"[INFO] No applications found for: %s\n",
			query,
		)

		lastExitCode = 0
		return
	}

	fmt.Printf(
		"Search results for \"%s\":\n\n",
		query,
	)

	for _, app := range results {
		fmt.Printf(
			"%s  %s  [%s]\n",
			app.Name,
			app.Version,
			app.ID,
		)
	}

	lastExitCode = 0
}

func showAppInfo(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: info <app>")
		lastExitCode = 1
		return
	}

	if err := packages.Info(args[0]); err != nil {
		fmt.Println("[ERROR]", err)
		lastExitCode = 1
		return
	}

	lastExitCode = 0
}
