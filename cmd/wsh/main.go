package main

import (
	"bufio"
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

const version = "0.6.0"

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

	case "help", "?":
		printHelp()

	case "version":
		fmt.Println("WinDroid Shell", version)

	case "clear", "cls":
		clearScreen()

	case "exit", "quit":
		fmt.Println("Goodbye.")
		return false

	case "pwd", "gl":
		printWorkingDirectory()

	case "cd":
		changeDirectory(args)

	case "dir", "ls", "gci":
		listDirectory(args)

	case "cat", "type", "gc":
		showFile(args)

	case "mkdir", "newdir":
		makeDirectory(args)

	case "new":
		newItem(args)

	case "del", "rm", "remove":
		deleteFile(args)

	case "cp", "copy":
		copyItem(args)

	case "mv", "move":
		moveItem(args)

	case "write":
		writeFile(args)

	case "append":
		appendFile(args)

	case "test":
		testPath(args)

	case "ps", "process":
		listProcesses()

	case "kill":
		killProcess(args)

	case "env":
		showEnvironment(args)

	case "date":
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	case "which":
		whichCommand(args)

	case "commands":
		listCommands()

	case "history":
		fmt.Println("[INFO] Command history is not persistent yet.")

	case "echo":
		fmt.Println(strings.Join(args, " "))

	case "set":
		setVariable(args)

	case "unset":
		unsetVariable(args)

	case "run":
		runApplication(args)

	case "wqa":
		runWQA(args)

	case "install":
		if len(args) == 0 {
			fmt.Println("Usage: install <package|app>")
			break
		}

		if err := installApplication(args[0]); err != nil {
			fmt.Println("[ERROR]", err)
		}

	case "update":
		if len(args) == 0 {
			fmt.Println("Usage: update <app|all>")
			break
		}

		if strings.EqualFold(args[0], "all") {
			if err := updater.UpdateAll(); err != nil {
				fmt.Println("[ERROR]", err)
			}
			break
		}

		if err := updater.Update(args[0]); err != nil {
			fmt.Println("[ERROR]", err)
		}

	case "remove-app":
		if len(args) == 0 {
			fmt.Println("Usage: remove-app <app>")
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
		showAppInfo(args)

	case "search":
		searchRepository(args)

	case "alias":
		handleAlias(args)

	default:
		runExternal(command, args)
	}

	return true
}

func printHelp() {
	fmt.Println(`
WinDroid Shell 0.6.0

File commands:

  pwd, gl
      Show current directory

  cd <path>
      Change directory

  dir, ls, gci
      List files and directories

  cat, type, gc <file>
      Show file contents

  mkdir, newdir <name>
      Create directory

  new <file>
      Create empty file

  del, rm, remove <file>
      Delete file

  cp, copy <source> <destination>
      Copy file

  mv, move <source> <destination>
      Move file

  write <file> <text>
      Write text to file

  append <file> <text>
      Append text to file

  test <path>
      Check whether path exists

System commands:

  ps, process
      List running processes

  kill <pid>
      Stop a process

  env
      Show environment variables

  env <name>
      Show one environment variable

  date
      Show current date and time

  which <command>
      Find command

  commands
      List available WSh commands

  history
      Show command history status

  echo <text>
      Print text

  set <name> <value>
      Set environment variable

  unset <name>
      Remove environment variable

Package commands:

  install <package|app>
      Install WQA package or repository application

  update <app>
      Update application

  update all
      Update all installed applications

  remove-app <app>
      Remove installed application

  list
      List installed applications

  info <app>
      Show application information

  search [name]
      Search WQA application repository

WQA:

  wqa <command>
      Run WQA CLI

Shell:

  help
      Show this help

  version
      Show WSh version

  clear
      Clear screen

  alias
      Show alias information

  exit
      Exit WSh

Run:

  run <program.exe>
      Run Windows executable`)
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
		} else {
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
	cmd := exec.Command(
		"tasklist",
	)

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

func listCommands() {
	fmt.Println(`
WSh commands:

help
version
clear
exit
pwd
cd
dir
ls
cat
type
mkdir
new
del
rm
cp
mv
write
append
test
ps
kill
env
date
which
commands
history
echo
set
unset
alias
run
wqa
install
update
remove-app
list
info
search`)
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

	fmt.Println("[OK] Environment variable set:", name)
}

func unsetVariable(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: unset <name>")
		return
	}

	if err := os.Unsetenv(args[0]); err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[OK] Environment variable removed:", args[0])
}

func handleAlias(args []string) {
	if len(args) == 0 {
		fmt.Println("WSh aliases:")
		fmt.Println("  ls       -> dir")
		fmt.Println("  gl       -> pwd")
		fmt.Println("  gc       -> cat")
		fmt.Println("  gci      -> dir")
		fmt.Println("  rm       -> del")
		fmt.Println("  cp       -> copy")
		fmt.Println("  mv       -> move")
		return
	}

	fmt.Println("[INFO] Alias management will be expanded in a future version.")
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
	exe, err := os.Executable()

	if err == nil {
		dir := filepath.Dir(exe)

		local := filepath.Join(
			dir,
			"wqa.exe",
		)

		if _, err := os.Stat(local); err == nil {
			return local, nil
		}
	}

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
			fmt.Println("[ERROR]", err)
		}

		return
	}

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

	if strings.Contains(target, `\`) ||
		strings.Contains(target, `/`) {

		if _, err := os.Stat(target); err == nil {
			runEXEFile(target)
			return
		}
	}

	localPath := filepath.Join(
		".",
		target,
	)

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

	if resolved, err := exec.LookPath(target); err == nil {
		runEXEFile(resolved)
		return
	}

	fmt.Println("[ERROR] Application not found:", target)
}

func runEXEFile(target string) {
	resolved, err := exec.LookPath(target)

	if err == nil {
		target = resolved
	} else {
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

	repo, err := repository.Load(
		"repository.json",
	)

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

	if err := packages.Info(query); err == nil {
		return
	}

	repo, err := repository.Load(
		"repository.json",
	)

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	app := repo.Find(query)

	if app == nil {
		fmt.Println(
			"[ERROR] Application not found:",
			query,
		)
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
	ext := strings.ToLower(
		filepath.Ext(target),
	)

	if ext == ".wqa" ||
		ext == ".exe" {
		return packages.Install(target)
	}

	repo, err := repository.Load(
		"repository.json",
	)

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

	if !strings.EqualFold(
		app.Type,
		"wqa",
	) {
		return fmt.Errorf(
			"unsupported repository package type: %s",
			app.Type,
		)
	}

	tempDir, err := os.MkdirTemp(
		"",
		"wqa-install-*",
	)

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

	if err := installer.Install(
		packagePath,
	); err != nil {
		return fmt.Errorf(
			"installation failed: %w",
			err,
		)
	}

	return nil
}
