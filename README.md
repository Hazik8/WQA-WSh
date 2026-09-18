# WQA & WSh

**WQA (WinDroid Application)** is an application language, bytecode runtime, and package system created by **Win Studio** for the WinDroid ecosystem.

**WSh (WinDroid Shell)** is a command-line shell for managing and running WQA applications and Windows programs.

## Versions

* **WQA:** 1.1.0
* **WSh:** 0.7.0
* **Runtime:** WQBC
* **Architectures:** x86_64, arm64

## WQA 1.1.0

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
wet       variable
print     output

add       addition
sub       subtraction
mul       multiplication
div       division

wif       condition
else      alternative branch
endwif    end condition

wloop     loop
endloop   end loop

repeat    repeat block
endrepeat end repeat

wfunc     function
endfunc   end function
call      call function
give      return value

winput    user input
wclear    clear console
wait      delay
wtime     current time
wdate     current date
wexit     exit program
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

## WQBC

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

## WSh

WSh provides command-line management for the WinDroid application ecosystem.

Example commands:

```text
help
version
clear
cd
pwd
dir
ls
type
cat
mkdir
del
rm
run
install
remove
list
search
wqa
exit
```

Example:

```text
wsh [C:\]> run Calculator.wqa
```

## Documentation

The complete WQA language specification is available in:

```text
WQA_LANGUAGE.md
```

Documentation repository:

`WQA-WSh-Docs`

## Project Structure

```text
WQA/
├── cmd/
│   ├── wqa/
│   └── wsh/
├── internal/
│   ├── compiler/
│   ├── format/
│   ├── lexer/
│   ├── parser/
│   ├── runtime/
│   └── ...
├── examples/
├── README.md
├── WQA_LANGUAGE.md
├── go.mod
└── LICENSE
```

## Project Status

WQA and WSh are actively developed as part of the WinDroid ecosystem.

**Current WQA version: 1.1.0**

**Current WSh version: 0.7.0**

## License

This project is licensed under the MIT License.

See `LICENSE` for details.

---

# Win Studio

WQA and WSh are projects of **Win Studio**.
