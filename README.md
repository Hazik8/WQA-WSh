# WQA & WSh

WQA (WinDroid Application) is an application format and runtime ecosystem created by Win Studio.

WSh (WinDroid Shell) is a lightweight command-line shell for working with WQA applications, Windows executables and installed applications.

## Features

- WSh command-line shell
- WQA application format
- WQA runtime
- WQA package installation
- Application management
- `run` command for `.wqa` and `.exe`
- Installed application management
- Local application search
- Windows executable support

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
Running applications

Run a Windows executable:

wsh> run notepad.exe

Run an executable by path:

wsh> run C:\Apps\TestApp.exe

Run an installed application:

wsh> run TestApp

Run a WQA application:

wsh> run Calculator.wqa
WQA

WQA applications use the .wqa format.

## Examples

Calculator.wqa

WQA applications contain a manifest and application data that can be executed by the WQA runtime.

Project status

Example WQA programs are available in the `examples` directory.

### Hello World

examples/hello/main.wqa

Calculator

examples/calculator/main.wqa

Conditions

examples/conditions/main.wqa

These examples demonstrate basic WQA output, variables, arithmetic and conditions.

WQA and WSh are currently under active development.

Some planned features are not implemented yet, including:

Online application repository
Automatic WSh updates
Application signatures
More WQA APIs
Cross-platform support
Building

Requirements:

Go
Windows

Build WSh:

go build -o wsh.exe .\cmd\wsh

Build WQA:

go build -o wqa.exe .\cmd\wqa
Project structure
WQA/
├── cmd/
│   ├── wsh/
│   └── wqa/
├── internal/
│   ├── packages/
│   ├── runtime/
│   └── format/
├── examples/
│   ├── calculator/
│   ├── conditions/
│   └── hello/
├── go.mod
├── README.md
└── LICENSE
License

This project is licensed under the MIT License.

See LICENSE for details.

Win Studio

WQA and WSh are projects of Win Studio