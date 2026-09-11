# WQA & WSh

WQA (WinDroid Application) is an application format and runtime ecosystem created by Win Studio.

WSh (WinDroid Shell) is a lightweight command-line shell for working with WQA applications, Windows executables, and installed applications.

## Versions

* WQA: 1.0.0
* WSh: 0.6.0

## Features

* WSh command-line shell
* WQA application format
* WQA runtime
* WQA language
* WQBC bytecode runtime
* WQA package installation
* Application management
* Run `.wqa` and `.exe` applications
* Installed application management
* Local application search
* Windows executable support
* WQA application updates
* Application update support

## WSh Commands

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

### Running Applications

Run a Windows executable:

```text
wsh> run notepad.exe
```

Run an executable by path:

```text
wsh> run C:\Apps\TestApp.exe
```

Run an installed application:

```text
wsh> run TestApp
```

Run a WQA application:

```text
wsh> run Calculator.wqa
```

### Installing Applications

Install a WQA package:

```text
wsh> install Calculator.wqa
```

WQA applications are installed into:

```text
C:\WinDroid\Apps
```

## WQA

WQA (WinDroid Application) is the application format used by the WinDroid application ecosystem.

WQA applications use the `.wqa` package format.

A WQA package can contain:

* WQA package header
* Application manifest
* Application data
* WQBC bytecode
* Resources

The WQA runtime loads the package and executes its WQBC bytecode.

## WQA Source Files

WQA source files use the `.wq` extension.

Every WQA source file must begin with the `wqa` header:

```wq
wqa

print "Hello, World!"
```

The `wqa` header identifies the file as WQA source code.

## WQA Language

WQA uses a simple command-based syntax.

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

### Variables

The `wet` keyword creates a variable:

```wq
wet x = 10
```

### Output

The `print` command outputs a value:

```wq
print "Hello, World!"
```

### Arithmetic

WQA supports basic arithmetic operations:

```wq
wet a = 10
wet b = 5

add a b result
sub a b difference
mul a b product
div a b quotient
```

### Conditions

Conditions use `wif`, `else`, and `endwif`:

```wq
wif x > 5
    print "YES"
else
    print "NO"
endwif
```

Supported comparison operators include:

```text
>
<
==
```

For the complete WQA language specification, see:

`WQA_LANGUAGE.md`

## WQA Projects

A WQA project contains a `wqa.json` manifest and a WQA source entry file.

Example:

```text
MyProject/
├── wqa.json
└── main.wq
```

Example `wqa.json`:

```json
{
    "name": "MyProject",
    "version": "1.0.0",
    "entry": "main.wq",
    "runtime": "wqbc"
}
```

## Building WQA Applications

Create a WQA project and place the source code and `wqa.json` manifest inside it.

Build the project with:

```text
wqa build .
```

The command produces a `.wqa` package.

Example:

```text
MyProject.wqa
```

Run the resulting package with:

```text
wqa run .\MyProject.wqa
```

## WQA Examples

Example WQA programs are available in the `examples` directory.

### Hello World

```text
examples/hello/
```

### Calculator

```text
examples/calculator/
```

### Conditions

```text
examples/conditions/
```

These examples demonstrate:

* Output
* Variables
* Arithmetic
* Conditions
* WQA project structure
* WQA package building

## WQBC

WQBC is the bytecode format used by the WQA runtime.

The general execution flow is:

```text
.wq source
    ↓
Lexer
    ↓
Parser
    ↓
Compiler
    ↓
WQBC bytecode
    ↓
.wqa package
    ↓
WQA Runtime
```

The runtime loads the compiled bytecode from the WQA package and executes it.

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
│   ├── packages/
│   ├── parser/
│   ├── runtime/
│   └── logger/
├── examples/
│   ├── calculator/
│   ├── conditions/
│   └── hello/
├── go.mod
├── README.md
├── WQA_LANGUAGE.md
└── LICENSE
```

## Building

### Requirements

* Go
* Windows

### Build WSh

```powershell
go build -o wsh.exe .\cmd\wsh
```

### Build WQA

```powershell
go build -o wqa.exe .\cmd\wqa
```

### Format the source code

```powershell
go fmt ./...
```

## Installation

WQA and WSh can be installed manually to locations such as:

```text
C:\WinDroid\Tools
C:\WinDroid\Bin
```

WSh uses:

```text
C:\WinDroid\Apps
```

as the default application installation directory.

The project also includes update/build tooling for building and installing newer versions.

## WSh Application Management

WSh can work with both WQA applications and Windows executables.

Examples:

```text
wsh> list
wsh> info Calculator
wsh> run Calculator.wqa
wsh> run notepad.exe
wsh> install Calculator.wqa
wsh> remove Calculator
```

WSh automatically works with installed applications and can also launch Windows executables available on the system.

## Project Status

WQA and WSh are currently under active development.

The WQA 1.0 language and package system are being developed alongside the WSh runtime environment.

## Planned Features

* Online application repository
* Automatic WSh updates
* Application signatures
* More WQA APIs
* Cross-platform support
* Improved package management
* Additional WQA language features
* Expanded WQBC instruction set

## License

This project is licensed under the MIT License.

See `LICENSE` for details.

## Win Studio

WQA and WSh are projects of Win Studio.
