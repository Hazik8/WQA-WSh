package main

import "fmt"

func printHelp() {
	fmt.Println(`
╔══════════════════════════════════════╗
║        WinDroid Shell 1.0.0          ║
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
      Delete file or directory

  cp, copy <source> <destination>
      Copy file or directory

  mv, move <source> <destination>
      Move file or directory

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

  help, ?
      Show this help

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

  version
      Show WSh version

  --version
      Show WSh version

  -v
      Show WSh version

  clear, cls
      Clear screen

  exit, quit
      Exit WSh


Run:

  run <program>
      Run Windows executable

  wqa <command>
      Run WQA CLI


Packages:

  repo
      Show repository usage

  repo update
      Update the application repository

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
      Search WQA application repository`)
}
