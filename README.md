# WQA & WSh

WQA (WinDroid Application) is an application format and runtime ecosystem created by Win Studio.

WSh (WinDroid Shell) is a lightweight command-line shell for working with WQA applications, Windows executables and installed applications.

## Versions

- WQA: 1.0.0
- WSh: 0.6.0

## Features

- WSh command-line shell
- WQA application format
- WQA runtime
- WQA package installation
- Application management
- Run `.wqa` and `.exe` applications
- Installed application management
- Local application search
- Windows executable support
- WQA application updates
- Application update support

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
Running Applications

Run a Windows executable:

wsh> run notepad.exe

Run an executable by path:

wsh> run C:\Apps\TestApp.exe

Run an installed application:

wsh> run TestApp

Run a WQA application:

wsh> run Calculator.wqa
Installing Applications

Install a WQA package:

wsh> install Calculator.wqa

WQA applications are installed into:

C:\WinDroid\Apps
WQA

WQA (WinDroid Application) is the application format used by the WinDroid application ecosystem.

WQA applications use the .wqa file format.

A WQA package contains:

WQA header
Application manifest
Application data
WQBC bytecode

The WQA runtime executes WQBC bytecode.

WQA Examples

Example WQA programs are available in the examples directory.

Hello World
examples/hello/main.wqa
Calculator
examples/calculator/main.wqa
Conditions
examples/conditions/main.wqa

These examples demonstrate:

Output
Variables
Arithmetic
Conditions
WQA Language

Example:

wet x = 10

wif x > 5
    print "YES"
else
    print "NO"
endwif

The wet keyword is used to create variables.

Project Structure
WQA/
├── cmd/
│   ├── wqa/
│   └── wsh/
├── internal/
│   ├── format/
│   ├── packages/
│   └── runtime/
├── examples/
│   ├── calculator/
│   ├── conditions/
│   └── hello/
├── go.mod
├── README.md
└── LICENSE
Building
Requirements
Go
Windows
Build WSh
go build -o wsh.exe .\cmd\wsh
Build WQA
go build -o wqa.exe .\cmd\wqa
Installation

WQA and WSh can be installed manually to:

C:\WinDroid\Tools
C:\WinDroid\Bin

The project also includes an update script for building and installing the latest versions.

Project Status

WQA and WSh are currently under active development.

Planned Features
Online application repository
Automatic WSh updates
Application signatures
More WQA APIs
Cross-platform support
Improved package management
License

This project is licensed under the MIT License.

See LICENSE for details.

Win Studio

WQA and WSh are projects of Win Studio.
