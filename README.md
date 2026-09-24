# WQA & WSh

**WQA (Win Quick Application)** is an application language, bytecode runtime, and package system created by **Win Studio** for the WinDroid ecosystem.

**WSh (Win Shell)** is a command-line shell for managing and running WQA applications and Windows programs.

## Versions

* **WQA:** 1.1.0
* **WSh:** 1.0.0
* **Runtime:** WQBC
* **Architectures:** x86_64, arm64

---

# WQA 1.1.0

WQA 1.1.0 expands the core language with:

* Variables with `wet`
* Output with `print`
* Arithmetic operations
* Conditions
* Loops
* Repeat blocks
* Functions
* Function calls
* Function return values
* User input
* Console clearing
* Delays
* Current time
* Current date
* Program exit

### Example

```wq
wqa

wet name = "WinDroid"

print "Hello,"
print name

winput answer

print "Your answer:"
print answer

wtime
wdate
```

## File Types

| Extension | Purpose                          |
| --------- | -------------------------------- |
| `.wq`     | WQA source code                  |
| `.wqa`    | Compiled WQA application package |
| `.wqbc`   | WQBC bytecode                    |

## WQA Source

Every `.wq` source file must begin with:

```wq
wqa
```

Example:

```wq
wqa

wet x = 10

wif x > 5
    print "YES"
else
    print "NO"
endwif
```

## Core Commands

```text
wet        variable
print      output
add        addition
sub        subtraction
mul        multiplication
div        division

wif        condition
else       alternative branch
endwif     end condition

wloop      loop
endloop    end loop

repeat     repeat block
endrepeat  end repeat

wfunc      function
endfunc    end function
call       call function
give       return value

winput     user input
wclear     clear console
wait       delay
wtime      current time
wdate      current date
wexit      exit program
```

## Building

Build a WQA project:

```powershell
wqa build .
```

Build a project from another directory:

```powershell
wqa build .\TestProject
```

The compiler produces:

```text
TestProject.wqa
```

Run a WQA package:

```powershell
wqa run .\TestProject.wqa
```

## Functions

```wq
wqa

wfunc hello
    print "Hello from WQA!"
endfunc

call hello
```

Functions can return values:

```wq
wqa

wfunc getNumber
    wet number = 42
    give number
endfunc

call getNumber result

print result
```

## Input

```wq
wqa

winput name

print "Hello:"
print name
```

## Time and Date

```wq
wqa

print "Current time:"
wtime

print "Current date:"
wdate
```

Example output:

```text
Current time:
23:24:31

Current date:
18.09.2026
```

---

# WQBC

WQBC is the bytecode runtime used to execute compiled WQA applications.

The execution pipeline is:

```text
.wq
  ↓
Lexer
  ↓
Parser
  ↓
Compiler
  ↓
WQBC
  ↓
.wqa
  ↓
WQBC Runtime
  ↓
Application
```

---

# WSh 1.0.0

WSh provides command-line management for the WinDroid application ecosystem.

WSh can:

* Manage files and directories
* Run Windows programs
* Run WQA applications
* Manage installed applications
* Search the WQA application repository
* Install and update applications
* Manage shell history
* Create command aliases
* Manage environment variables

## WSh Commands

### Shell

```text
help, ?
version
--version
-v
clear, cls
history
alias
unalias
exit, quit
```

### Files

```text
pwd, gl
cd <path>
dir, ls, gci
cat, type, gc <file>
mkdir, newdir <name>
new <file>
del, rm, remove <file>
cp, copy <source> <destination>
mv, move <source> <destination>
write <file> <text>
append <file> <text>
test <path>
```

### System

```text
ps, process
kill <pid>
env
env <name>
set <name> <value>
unset <name>
date
which <command>
echo <text>
```

### Run

```text
run <program>
wqa <command>
```

### Packages

```text
repo
repo update
install <package|app>
update <app>
update all
remove-app <app>
list
info <app>
search [name]
```

## WSh Example

```text
wsh [C:\WQA]> wqa build
[OK] Created: TestProject.wqa

wsh [C:\WQA]> wqa run TestProject.wqa
[INFO] Loading: TestProject.wqa
[INFO] Starting WQBC Runtime
```

WSh can also run regular Windows programs:

```text
wsh [C:\WQA]> run notepad.exe
[OK] Process started: C:\Windows\system32\notepad.exe
```

## WSh Application Repository

WSh includes an application repository system for WQA applications.

Repository operations include:

```text
repo update
search Calculator
info windroid.Calculator
install windroid.Calculator
list
```

Applications are downloaded and verified using SHA-256 before installation.

---

# Documentation

The complete WQA language specification is available in:

[WQA_LANGUAGE.md](https://github.com/Hazik8/WQA-WSh-Docs/WQA_LANGUAGE.md)

Documentation repository:

[WQA-WSh-Docs](https://github.com/Hazik8/WQA-WSh-Docs)

---

# Project Structure

```text
WQA/

├── cmd/
│   ├── wqa/
│   └── wsh/
│
├── internal/
│   ├── compiler/
│   ├── format/
│   ├── lexer/
│   ├── parser/
│   ├── runtime/
│   └── ...
│
├── examples/
├── README.md
├── WQA_LANGUAGE.md
├── go.mod
└── LICENSE
```

---

# Project Status

WQA and WSh are actively developed as part of the WinDroid ecosystem.

**Current WQA version: 1.1.0**

**Current WSh version: 1.0.0**

WSh 1.0.0 focuses on a stable command-line environment for Windows and integration with the WQA application ecosystem.

---

# License

This project is licensed under the MIT License.

See `LICENSE` for details.

---

# Win Stydio

WQA and WSh are projects of **Win Stydio**.
