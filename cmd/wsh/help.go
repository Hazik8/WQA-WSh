package main

import "fmt"

func printHelp() {
	fmt.Println(`
╔══════════════════════════════════════╗
║        WinDroid Shell 0.7.0         ║
╚══════════════════════════════════════╝

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

  repo
      Application repository management
      
  repo update
      Update the local application repository


System commands:

  ps, process
      List running processes

  kill <pid>
      Stop a process

  env
      Show environment variables

  env <name>
      Show one environment variable

  set <name> <value>
      Set environment variable

  unset <name>
      Remove environment variable

  date
      Show current date and time

  which <command>
      Find command

  echo <text>
      Print text


Shell:

  history
      Show command history

  history <number>
      Show last N commands

  alias
      Show aliases

  alias <name> <command>
      Create an alias

  unalias <name>
      Remove an alias

  commands
      List commands

  version
      Show WSh version

  --version
      Show WSh version

  clear, cls
      Clear screen

  exit, quit
      Exit WSh


Run:

  run <program.exe>
      Run Windows executable


Packages:

  install <package|app>
      Install application

  update <app>
      Update application

  update all
      Update all applications

  remove-app <app>
      Remove application

  list
      List installed applications

  info <app>
      Show application information

  search [name]
      Search WQA application repository


WQA:

  wqa <command>
      Run WQA CLI`)
}
