package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const version = "0.7.0"

var (
	commandHistory []string
	aliases        = make(map[string]string)
	lastExitCode   = 0
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	loadHistory()
	loadAliases()

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

		addHistory(input)

		if !handleCommand(input) {
			return
		}
	}
}

func handleCommand(input string) bool {
	parts := parseCommandLine(input)

	if len(parts) == 0 {
		return true
	}

	command := strings.ToLower(parts[0])
	args := parts[1:]

	// Shell aliases
	if aliasValue, ok := aliases[command]; ok {
		expanded := aliasValue

		if len(args) > 0 {
			expanded += " " + strings.Join(args, " ")
		}

		return handleCommand(expanded)
	}

	switch command {
	case "help", "?":
		printHelp()

	case "version", "--version", "-v":
		fmt.Println("WinDroid Shell", version)

	case "clear", "cls":
		clearScreen()

	case "exit", "quit":
		fmt.Println("Goodbye.")
		saveHistory()
		saveAliases()
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
		showDate()

	case "which":
		whichCommand(args)

	case "history":
		showHistory(args)

	case "echo":
		echoCommand(args)

	case "set":
		setVariable(args)

	case "unset":
		unsetVariable(args)

	case "run":
		runApplication(args)

	case "wqa":
		runWQA(args)

	case "install":
		handleInstall(args)

	case "update":
		handleUpdate(args)

	case "remove-app":
		handleRemoveApp(args)

	case "repo":
		handleRepo(args)

	case "list":
		handlePackageList()

	case "info":
		showAppInfo(args)

	case "search":
		searchRepository(args)

	case "alias":
		handleAlias(args)

	case "unalias":
		handleUnalias(args)

	default:
		runExternal(command, args)
	}

	return true
}
