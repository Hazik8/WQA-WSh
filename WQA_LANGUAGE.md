# WQA Language Specification

**WQA Language Version:** 1.0
**Runtime:** WQBC
**Architecture:** x86_64, arm64
**Project:** WQA / WSh
**Publisher:** Win Studio
**Status:** Early Development Specification

---

## 1. About WQA

WQA is the application programming language used by the WinDroid application ecosystem.

WQA is designed to provide a small, simple, and predictable language for creating applications that can be compiled into `.wqa` application packages and executed by the WQA Runtime.

The language is intentionally minimal.

Instead of providing a large number of language features, WQA focuses on a small instruction set that can be implemented efficiently by the WQA compiler and WQBC virtual machine.

A basic WQA program looks like this:

```wqa
wqa

print "Hello, WQA!"
```

The `wqa` header identifies the source file as a WQA program.

The header must be written in lowercase.

---

## 2. WQA Source Header

Every WQA source file must begin with the following line:

```wqa
wqa
```

The header is mandatory.

Correct:

```wqa
wqa

print "Hello, WQA!"
```

Incorrect:

```text
WQA
```

Incorrect:

```text
Wqa
```

The WQA language specification defines the source header as the exact lowercase keyword:

```text
wqa
```

The header must appear at the beginning of the source file, before any WQA instructions.

### 2.1 Why the Header Exists

The header provides an explicit indication that a file contains WQA source code.

It also allows tools such as the WQA compiler to validate the source format before compilation begins.

For example, the WQA compiler can reject a source file that does not begin with:

```text
wqa
```

---

## 3. File Types

WQA uses several file extensions.

| Extension | Purpose                 |
| --------- | ----------------------- |
| `.wq`     | WQA source file         |
| `.wqbc`   | WQBC bytecode file      |
| `.wqa`    | WQA application package |

### 3.1 `.wq`

A `.wq` file contains human-readable WQA source code.

Example:

```wqa
wqa

wet score = 100
print score
```

Source files are intended to be compiled by the WQA compiler.

### 3.2 `.wqbc`

A `.wqbc` file contains WQBC bytecode.

WQBC is the bytecode representation executed by the WQA Runtime.

### 3.3 `.wqa`

A `.wqa` file is a packaged WQA application.

A WQA package can contain:

* Application metadata
* Manifest data
* WQBC bytecode
* Application resources
* Other package data

---

## 4. Hello World

The simplest WQA program is:

```wqa
wqa

print "Hello, WQA!"
```

Output:

```text
Hello, WQA!
```

The first line identifies the source as WQA.

The `print` instruction writes the specified value to the output.

---

## 5. Output

The `print` instruction is used to display values.

Example:

```wqa
wqa

print "Hello"
print "Welcome to WQA"
```

Output:

```text
Hello
Welcome to WQA
```

Each `print` instruction produces a new line.

### 5.1 Printing a String

```wqa
wqa

print "Hello, WinDroid!"
```

### 5.2 Printing a Variable

Variables can also be used as expressions:

```wqa
wqa

wet name = "WinDroid"

print name
```

The runtime resolves the variable and prints its value.

---

## 6. Variables

Variables store values that can be used by WQA instructions.

The current WQA syntax uses the `wet` keyword for variable assignment.

Basic syntax:

```text
wet name = value
```

Example:

```wqa
wqa

wet score = 100
wet lives = 3
wet name = "Calculator"
```

### 6.1 Variable Names

Variable names may contain:

* Letters
* Numbers
* Underscores

A variable name must not begin with a number.

Valid examples:

```text
score
player
result
lives
player_name
value1
```

Invalid example:

```text
123value
```

### 6.2 Updating Variables

A variable can be assigned a new value:

```wqa
wqa

wet score = 10
wet score = 20

print score
```

The latest value is stored in the variable.

---

## 7. The `wet` Instruction

`wet` creates or updates a variable.

Syntax:

```text
wet variable = expression
```

Example:

```wqa
wqa

wet score = 100
```

Another example:

```wqa
wqa

wet name = "WinDroid"
```

The value may be a number, string, or variable expression supported by the current runtime.

### 7.1 Number Assignment

```wqa
wqa

wet lives = 3
wet score = 100
```

### 7.2 String Assignment

```wqa
wqa

wet name = "WinDroid"
wet message = "Hello!"
```

### 7.3 Variable Assignment

```wqa
wqa

wet score = 100
wet result = score

print result
```

---

## 8. Numbers

WQA currently supports integer numeric values.

Example:

```wqa
wqa

wet score = 100
wet lives = 3
wet version = 1
```

Numbers can be used by arithmetic instructions and comparisons.

Example:

```wqa
wqa

wet a = 10
wet b = 20

add a b result

print result
```

Output:

```text
30
```

---

## 9. Strings

Strings are text values enclosed in double quotation marks.

Example:

```wqa
wqa

wet name = "WinDroid"
wet message = "Hello!"
```

Strings can be printed:

```wqa
wqa

wet message = "Hello, WQA!"

print message
```

Output:

```text
Hello, WQA!
```

### 9.1 String Syntax

A string begins and ends with:

```text
"
```

Example:

```wqa
"Hello"
```

Strings may contain spaces.

Example:

```wqa
"Hello, WinDroid!"
```

An unterminated string is a lexer error.

---

## 10. Expressions

An expression represents a value used by an instruction.

The current WQA parser supports:

* Identifiers
* Numbers
* Strings

Examples:

```text
score
100
"Hello"
```

Expressions can be used with instructions such as:

```wqa
print score
```

```wqa
print "Hello"
```

```wqa
wet result = score
```

---

## 11. Addition

The `add` instruction adds two values and stores the result in a variable.

Syntax:

```text
add left right result
```

Example:

```wqa
wqa

wet a = 10
wet b = 20

add a b result

print result
```

Output:

```text
30
```

The third argument specifies the variable where the result is stored.

### 11.1 Addition Example

```wqa
wqa

wet apples = 5
wet oranges = 7

add apples oranges total

print total
```

---

## 12. Subtraction

The `sub` instruction subtracts the second value from the first value.

Syntax:

```text
sub left right result
```

Example:

```wqa
wqa

wet a = 30
wet b = 10

sub a b result

print result
```

Output:

```text
20
```

The result is stored in the variable specified by the third argument.

---

## 13. Multiplication

The `mul` instruction multiplies two values.

Syntax:

```text
mul left right result
```

Example:

```wqa
wqa

wet a = 5
wet b = 4

mul a b result

print result
```

Output:

```text
20
```

---

## 14. Division

The `div` instruction divides the first value by the second value.

Syntax:

```text
div left right result
```

Example:

```wqa
wqa

wet a = 20
wet b = 5

div a b result

print result
```

Output:

```text
4
```

### 14.1 Division by Zero

Division by zero is not a valid operation.

Example:

```wqa
wqa

wet a = 20
wet b = 0

div a b result
```

A compliant runtime must not perform an invalid division by zero.

---

## 15. Arithmetic Instructions

The current arithmetic instruction set is:

| Instruction | Description    |
| ----------- | -------------- |
| `add`       | Addition       |
| `sub`       | Subtraction    |
| `mul`       | Multiplication |
| `div`       | Division       |

The standard form is:

```text
operation left right result
```

Example:

```wqa
wqa

wet a = 10
wet b = 5

add a b sum
sub a b difference
mul a b product
div a b quotient
```

---

## 16. Conditions

WQA supports conditional execution using the `wif` instruction.

Basic syntax:

```wqa
wif condition
    instructions
endwif
```

Example:

```wqa
wqa

wet score = 100

wif score > 50
    print "Winner"
endwif
```

If the condition evaluates to true, the instructions inside the block are executed.

If the condition evaluates to false, they are skipped.

---

## 17. Comparison Operators

The current WQA language supports the following comparison operators:

```text
>
<
==
```

### 17.1 Greater Than

The `>` operator checks whether the left value is greater than the right value.

Example:

```wqa
wqa

wet score = 100

wif score > 50
    print "Winner"
endwif
```

### 17.2 Less Than

The `<` operator checks whether the left value is less than the right value.

Example:

```wqa
wqa

wet score = 25

wif score < 50
    print "Try again"
endwif
```

### 17.3 Equal

The `==` operator checks whether two values are equal.

Example:

```wqa
wqa

wet score = 100

wif score == 100
    print "Perfect"
endwif
```

---

## 18. `else`

The `else` keyword provides an alternative branch for a `wif` statement.

Syntax:

```wqa
wif condition
    instructions
else
    instructions
endwif
```

Example:

```wqa
wqa

wet score = 40

wif score > 50
    print "Winner"
else
    print "Try again"
endwif
```

If the condition is true, the first block is executed.

If the condition is false, the `else` block is executed.

---

## 19. End of Condition

Every `wif` block must be closed with:

```text
endwif
```

Example:

```wqa
wqa

wet score = 100

wif score > 50
    print "Winner"
endwif
```

A `wif` statement without a corresponding `endwif` is invalid WQA syntax.

---

## 20. Nested Conditions

WQA condition blocks may be used inside other blocks when supported by the compiler and runtime.

Example:

```wqa
wqa

wet score = 100
wet lives = 3

wif score > 50
    print "Winner"

    wif lives > 0
        print "Player is active"
    endwif
endwif
```

Each `wif` must have its own `endwif`.

---

## 21. Complete Conditional Example

```wqa
wqa

wet score = 100

print "Game started"

wif score > 50
    print "Winner"
else
    print "Try again"
endwif

print "Game finished"
```

Expected output:

```text
Game started
Winner
Game finished
```

---

## 22. Comments

Comments are currently planned for a future version of WQA.

WQA 1.0 source code should not depend on comment syntax.

For maximum compatibility, source files should contain only syntax defined by this specification.

Future WQA versions may introduce a dedicated comment syntax.

---

## 23. Functions

User-defined functions are not part of WQA 1.0.

A possible future syntax could look like:

```text
function hello
    print "Hello!"
end
```

This syntax is only an example of a possible future feature.

It is **not valid WQA 1.0 syntax**.

---

## 24. WQBC

WQBC is the bytecode execution format used by the WQA Runtime.

WQA source code is processed by the compiler and converted into bytecode.

The general execution pipeline is:

```text
WQA Source
    |
    v
Lexer
    |
    v
Parser
    |
    v
Compiler
    |
    v
WQBC Bytecode
    |
    v
WQA Runtime
    |
    v
Program Output
```

The WQBC virtual machine processes bytecode instructions and maintains runtime state such as variables and the current instruction position.

---

## 25. WQBC Instruction Set

The current WQBC instruction set includes operations corresponding to WQA language features.

| Instruction | Description                 |
| ----------- | --------------------------- |
| `PRINT`     | Print a value               |
| `SET`       | Create or update a variable |
| `LOAD`      | Load a variable value       |
| `ADD`       | Addition                    |
| `SUB`       | Subtraction                 |
| `MUL`       | Multiplication              |
| `DIV`       | Division                    |
| `IF`        | Conditional execution       |
| `ENDIF`     | End a conditional block     |
| `EXIT`      | Stop execution              |

The WQBC instruction names are runtime-level operations.

They do not necessarily correspond one-to-one with source-level syntax.

For example, the WQA source instruction:

```text
wet
```

may be represented internally by a WQBC `SET` instruction.

---

## 26. WQBC Variables

Variables are stored by name.

Examples:

```text
score
name
result
lives
```

A runtime variable can contain a numeric or string value.

Example source:

```wqa
wqa

wet score = 100
wet name = "WinDroid"

print score
print name
```

The runtime maintains the values while the program is executing.

---

## 27. WQA Compiler

The WQA compiler converts WQA source code into executable WQBC bytecode and packages it into a `.wqa` application package.

A simplified compilation process is:

```text
.wq
 |
 +--> Header validation
 |
 +--> Lexer
 |
 +--> Parser
 |
 +--> Compiler
 |
 +--> WQBC
 |
 +--> Manifest
 |
 +--> WQA Package
```

The compiler is responsible for detecting invalid source syntax before the application is packaged.

---

## 28. WQA Project Structure

A basic WQA project may contain a structure similar to:

```text
MyApplication/
│
├── wqa.json
│
└── app/
    └── main.wq
```

The manifest identifies the application's entry point.

Example:

```json
{
  "wqa": "1.0",
  "id": "windroid.MyApplication",
  "name": "MyApplication",
  "version": "1.0.0",
  "publisher": "Win Studio",
  "runtime": "wqbc",
  "entry": "app/main.wq",
  "architecture": [
    "x86_64",
    "arm64"
  ]
}
```

---

## 29. WQA Application Manifest

A WQA application package contains a manifest describing the application.

Example:

```json
{
  "wqa": "1.0",
  "id": "windroid.Calculator",
  "name": "Calculator",
  "version": "1.0.0",
  "publisher": "Win Studio",
  "runtime": "wqbc",
  "entry": "app/main.wq",
  "architecture": [
    "x86_64",
    "arm64"
  ]
}
```

---

## 30. Manifest Fields

| Field          | Description                     |
| -------------- | ------------------------------- |
| `wqa`          | WQA package format version      |
| `id`           | Unique application identifier   |
| `name`         | Application name                |
| `version`      | Application version             |
| `publisher`    | Application publisher           |
| `runtime`      | Runtime used by the application |
| `entry`        | Application source entry point  |
| `architecture` | Supported CPU architectures     |

### 30.1 `wqa`

Defines the WQA package format version.

Current value:

```json
"wqa": "1.0"
```

### 30.2 `id`

Defines the unique application identifier.

Example:

```json
"id": "windroid.Calculator"
```

### 30.3 `name`

Defines the human-readable application name.

Example:

```json
"name": "Calculator"
```

### 30.4 `version`

Defines the application version.

Example:

```json
"version": "1.0.0"
```

### 30.5 `publisher`

Defines the publisher of the application.

Example:

```json
"publisher": "Win Studio"
```

### 30.6 `runtime`

Defines the runtime used by the application.

Current WQA 1.0 applications use:

```json
"runtime": "wqbc"
```

### 30.7 `entry`

Defines the application's entry source file.

Example:

```json
"entry": "app/main.wq"
```

### 30.8 `architecture`

Defines the CPU architectures supported by the application.

Example:

```json
"architecture": [
  "x86_64",
  "arm64"
]
```

---

## 31. Application IDs

Every installed WQA application should have a unique application ID.

Example:

```text
windroid.Calculator
```

Application IDs are used to identify applications within the WinDroid ecosystem.

The WSh package manager can use the application ID when locating or launching installed applications.

Example:

```text
wsh> run windroid.Calculator
```

The exact command behavior depends on the WSh implementation.

---

## 32. WQA Packages

A WQA application can be distributed as:

```text
Calculator.wqa
```

The `.wqa` package is a binary application container.

A package can contain:

```text
Manifest
WQBC bytecode
Resources
Application data
```

The package is designed to allow applications to be distributed as a single file.

---

## 33. WQA Binary Package Header

The WQA package format contains a binary header.

The header identifies the package and provides information about its internal sections.

The current format includes fields for:

```text
Magic
Version
Header Size
Flags
Manifest Offset
Manifest Size
Application Offset
Application Size
Resources Offset
Resources Size
Signature
CRC32
```

The package header is currently designed around a 64-byte header.

---

## 34. WQA Magic

The WQA binary package begins with the WQA magic value:

```text
WQA\x00
```

This value identifies the file as a WQA package.

The binary package magic is different from the source-language header.

### Source Header

```text
wqa
```

### Binary Package Magic

```text
WQA\x00
```

The source header is part of the WQA language.

The binary magic is part of the `.wqa` package format.

---

## 35. WQA Package Version

The current WQA package format version is:

```text
1.0
```

The binary format stores major and minor version information.

Current values:

```text
Major Version: 1
Minor Version: 0
```

Future versions may introduce additional package features.

---

## 36. WQA Package Flags

The package format defines flags for optional package features.

Current flag types include:

```text
Compressed
Signed
Encrypted
Executable
```

These flags allow the package format to describe additional properties of an application package.

Support for individual features depends on the current WQA runtime and package implementation.

---

## 37. Package Sections

A WQA package is divided into logical sections.

The main sections are:

```text
Header
Manifest
Application
Resources
```

The header stores offsets and sizes that allow the runtime to locate these sections.

A simplified package layout is:

```text
+----------------------+
| WQA Header           |
+----------------------+
| Manifest             |
+----------------------+
| WQBC Application     |
+----------------------+
| Resources            |
+----------------------+
```

---

## 38. Package Validation

Before executing a WQA package, the runtime should validate package information.

Validation may include:

* Magic value
* Package version
* Header size
* Section offsets
* Section sizes
* Manifest data
* CRC32
* Signature information when supported

Invalid packages should not be executed as valid WQA applications.

---

## 39. Runtime

WQA programs are executed by the WQA Runtime.

The runtime provides the execution environment for WQBC applications.

Current runtime responsibilities include:

* Loading WQA packages
* Reading package metadata
* Loading WQBC bytecode
* Managing variables
* Performing arithmetic
* Evaluating conditions
* Producing output
* Handling program execution

The runtime is responsible for interpreting or executing WQBC instructions.

---

## 40. Virtual Machine

The WQBC runtime contains a virtual machine.

The virtual machine maintains execution state such as:

```text
Instruction Pointer
Bytecode
Variables
```

A simplified runtime model is:

```text
VM
├── Code
├── IP
└── Variables
```

The instruction pointer identifies the current position in the bytecode.

The variable storage contains values created during execution.

---

## 41. Runtime Execution

A WQA application generally follows this execution sequence:

```text
Load package
      |
      v
Validate header
      |
      v
Read manifest
      |
      v
Load bytecode
      |
      v
Start WQBC VM
      |
      v
Execute instructions
      |
      v
Produce output
```

The runtime should reject invalid or incompatible package data.

---

## 42. Architecture

WQA applications can declare the CPU architectures they support.

Current architecture identifiers are:

```text
x86_64
arm64
```

These identifiers are declared in the application manifest.

Example:

```json
"architecture": [
  "x86_64",
  "arm64"
]
```

### 42.1 x86_64

`x86_64` represents 64-bit x86 processors.

### 42.2 arm64

`arm64` represents 64-bit ARM processors.

---

## 43. Versioning

WQA applications use semantic versioning.

Example:

```text
1.0.0
```

The format is:

```text
MAJOR.MINOR.PATCH
```

### MAJOR

The major version changes when incompatible application changes are introduced.

### MINOR

The minor version changes when compatible functionality is added.

### PATCH

The patch version changes for compatible fixes and corrections.

---

## 44. WQA Format Version

The WQA language and package format currently use:

```text
WQA 1.0
```

The application version and WQA format version are separate concepts.

For example:

```json
{
  "wqa": "1.0",
  "version": "2.4.1"
}
```

Here:

* `wqa` is the package format version.
* `version` is the application version.

---

## 45. Compatibility

A WQA 1.0 runtime should support the language and package features defined by this specification.

Applications should avoid depending on undocumented behavior.

Applications should declare their WQA version in the package manifest.

Example:

```json
"wqa": "1.0"
```

If a package requires a newer WQA version than the installed runtime supports, the runtime should reject the package rather than executing it incorrectly.

---

## 46. WQA and WSh

WQA and WSh are designed to work together within the WinDroid ecosystem.

WQA provides:

```text
Language
Compiler
Bytecode
Application Packages
Runtime
```

WSh provides:

```text
Shell
Application Management
Application Launching
Package Management
System Interaction
```

A WQA application can be installed and launched through WSh.

Example:

```text
wsh> install Calculator\Calculator.wqa
```

After installation, the application can be managed by WSh.

---

## 47. Building a WQA Application

A WQA project can be compiled with:

```text
wqa build .
```

The command should be executed from the WQA project directory.

Example:

```text
C:\Projects\Calculator> wqa build .
```

A successful build produces a `.wqa` package.

Example:

```text
Calculator.wqa
```

---

## 48. Running a WQA Package

A WQA package can be executed through the WQA command-line tool.

Example:

```text
wqa run .\Calculator.wqa
```

The runtime loads the package, reads its manifest, loads the WQBC bytecode, and starts execution.

---

## 49. Minimal Calculator

The following program demonstrates variables, arithmetic, and output:

```wqa
wqa

wet a = 10
wet b = 5

add a b result

print "Result:"
print result
```

Expected output:

```text
Result:
15
```

---

## 50. Complete Example

The following example demonstrates several core WQA features:

```wqa
wqa

wet score = 100

print "Game started"

wif score > 50
    print "Winner"
else
    print "Try again"
endwif

print "Game finished"
```

Expected output:

```text
Game started
Winner
Game finished
```

---

## 51. Example Application

A simple application project may look like:

```text
Game/
│
├── wqa.json
│
└── app/
    └── main.wq
```

`main.wq`:

```wqa
wqa

wet score = 100
wet lives = 3

print "Game started"

wif score > 50
    print "Winner"
else
    print "Game over"
endwif
```

`wqa.json`:

```json
{
  "wqa": "1.0",
  "id": "windroid.Game",
  "name": "Game",
  "version": "1.0.0",
  "publisher": "Win Studio",
  "runtime": "wqbc",
  "entry": "app/main.wq",
  "architecture": [
    "x86_64",
    "arm64"
  ]
}
```

The project can then be built with:

```text
wqa build .
```

---

## 52. Language Keywords

The current WQA source language defines the following keywords:

| Keyword  | Purpose                        |
| -------- | ------------------------------ |
| `wqa`    | WQA source header              |
| `wet`    | Create or update a variable    |
| `print`  | Output a value                 |
| `add`    | Add two values                 |
| `sub`    | Subtract two values            |
| `mul`    | Multiply two values            |
| `div`    | Divide two values              |
| `wif`    | Start a conditional block      |
| `else`   | Alternative conditional branch |
| `endwif` | End a conditional block        |

These keywords form the core WQA 1.0 source language.

---

## 53. Identifiers

Identifiers are names used for variables and other language objects.

Valid examples:

```text
score
player
result
lives
player_name
value1
```

Identifiers may contain:

```text
A-Z
a-z
0-9
_
```

An identifier cannot begin with a number.

---

## 54. Whitespace

Spaces and tabs can be used to separate WQA tokens.

Example:

```wqa
wet score = 100
```

Whitespace is also commonly used to make conditional blocks easier to read.

Example:

```wqa
wif score > 50
    print "Winner"
endwif
```

Indentation is primarily for readability.

The important structural tokens are:

```text
wif
else
endwif
```

---

## 55. New Lines

WQA uses new lines to separate instructions.

Example:

```wqa
wet a = 10
wet b = 20
print a
print b
```

Each instruction normally appears on its own line.

Empty lines may be used to improve readability.

Example:

```wqa
wqa

wet a = 10
wet b = 20

add a b result

print result
```

---

## 56. Error Handling

WQA tools can report errors during different stages of processing.

The major stages are:

```text
Lexer
Parser
Compiler
Package
Runtime
```

### 56.1 Lexer Errors

A lexer error occurs when the source contains an invalid character or malformed token.

Example:

```text
lexer error at line:column
```

### 56.2 Parser Errors

A parser error occurs when the token sequence does not follow the WQA grammar.

Example:

```text
parser error at line:column
```

### 56.3 Compiler Errors

A compiler error occurs when valid syntax cannot be converted into supported bytecode.

### 56.4 Runtime Errors

A runtime error occurs while executing a WQBC program.

Examples include invalid operations or invalid package data.

---

## 57. Source Validation

The WQA compiler should validate the source header before compilation.

A valid source file begins with:

```text
wqa
```

For example:

```wqa
wqa

print "Hello"
```

A file beginning with another value is not a valid WQA 1.0 source file.

---

## 58. Current Language Limitations

WQA 1.0 is intentionally small.

The current language does not define:

* User-defined functions
* Arrays
* Objects
* Loops
* Boolean literals
* Logical operators
* File APIs
* Networking APIs
* GUI APIs
* General-purpose system APIs
* Advanced string manipulation
* User-defined types

These features may be added in later versions.

---

## 59. Future Features

The following features may be introduced in future WQA versions:

* Functions
* Boolean values
* More comparison operators
* Logical operators
* Loops
* Arrays
* Improved string operations
* File APIs
* System APIs
* GUI APIs
* Networking APIs
* Application permissions
* Cryptographic signatures
* Package resources
* Cross-platform runtime
* Improved error handling
* Standard libraries

Future features are not guaranteed to be compatible with WQA 1.0.

---

## 60. Package Security

The WQA package format includes fields intended to support package security features.

These include:

```text
Signature
CRC32
Encrypted flag
Signed flag
```

Security features may be expanded in future WQA versions.

A runtime should not treat the existence of a package field as proof that the corresponding security feature is currently implemented.

---

## 61. Package Resources

The WQA package format reserves space for application resources.

Possible resources include:

```text
Images
Audio
Configuration files
Application data
Other assets
```

Resource support may be expanded in future versions of the WQA runtime.

Applications should not depend on undocumented resource behavior.

---

## 62. Cross-Platform Design

WQA is designed with multiple architectures in mind.

Current declared architectures:

```text
x86_64
arm64
```

The language itself is intended to remain independent from CPU architecture.

WQBC provides an intermediate execution layer between WQA applications and the underlying hardware.

This design allows the runtime to be implemented for multiple platforms.

---

## 63. Design Philosophy

WQA is designed to be:

* Small
* Fast
* Easy to learn
* Easy to implement
* Easy to package
* Suitable for application development
* Suitable for the WinDroid ecosystem

The language intentionally begins with a small instruction set.

Advanced functionality can be introduced through future language versions, runtime APIs, libraries, and package capabilities.

---

## 64. Language Stability

WQA Language 1.0 is an early development specification.

The language, compiler, bytecode format, package format, and runtime may change before the first stable release.

Developers should therefore:

* Declare the required WQA version
* Avoid undocumented features
* Avoid depending on implementation-specific behavior
* Keep application metadata up to date

---

## 65. Example: Basic Program

```wqa
wqa

print "Hello, WQA!"
```

---

## 66. Example: Variables

```wqa
wqa

wet name = "WinDroid"
wet version = 1

print name
print version
```

---

## 67. Example: Arithmetic

```wqa
wqa

wet a = 10
wet b = 5

add a b sum
sub a b difference
mul a b product
div a b quotient

print sum
print difference
print product
print quotient
```

---

## 68. Example: Conditions

```wqa
wqa

wet score = 100

wif score > 50
    print "Winner"
else
    print "Try again"
endwif
```

---

## 69. Example: Complete Calculator

```wqa
wqa

wet a = 10
wet b = 5

add a b sum
sub a b difference
mul a b product
div a b quotient

print "Addition:"
print sum

print "Subtraction:"
print difference

print "Multiplication:"
print product

print "Division:"
print quotient
```

Expected output:

```text
Addition:
15
Subtraction:
5
Multiplication:
50
Division:
2
```

---

## 70. Example: WinDroid Application

```wqa
wqa

wet app_name = "Calculator"
wet version = "1.0.0"

print app_name
print version

wet a = 20
wet b = 10

add a b result

wif result > 20
    print "Calculation complete"
else
    print "Calculation result is small"
endwif

print result
```

---

## 71. Recommended WQA Coding Style

WQA source code should be kept simple and readable.

Recommended:

```wqa
wqa

wet score = 100
wet lives = 3

wif score > 50
    print "Winner"
else
    print "Try again"
endwif
```

Avoid unnecessarily compressed code.

Use empty lines to separate logical sections.

Use indentation inside conditional blocks.

---

## 72. Minimal Valid WQA Program

The smallest valid WQA source file contains the WQA header:

```wqa
wqa
```

A useful executable program should contain at least one instruction.

Example:

```wqa
wqa

print "Hello"
```

---

## 73. WQA 1.0 Quick Reference

### Header

```text
wqa
```

### Variable

```text
wet name = value
```

### Output

```text
print value
```

### Addition

```text
add left right result
```

### Subtraction

```text
sub left right result
```

### Multiplication

```text
mul left right result
```

### Division

```text
div left right result
```

### Condition

```text
wif condition
    instructions
endwif
```

### Else

```text
wif condition
    instructions
else
    instructions
endwif
```

### Comparisons

```text
>
<
==
```

---

## 74. WQA 1.0 Keyword Reference

| Keyword  | Category     | Description                    |
| -------- | ------------ | ------------------------------ |
| `wqa`    | Header       | Identifies WQA source          |
| `wet`    | Variable     | Creates or updates a variable  |
| `print`  | Output       | Displays a value               |
| `add`    | Arithmetic   | Adds two values                |
| `sub`    | Arithmetic   | Subtracts two values           |
| `mul`    | Arithmetic   | Multiplies two values          |
| `div`    | Arithmetic   | Divides two values             |
| `wif`    | Control flow | Starts a conditional           |
| `else`   | Control flow | Defines the alternative branch |
| `endwif` | Control flow | Ends a conditional             |

---

## 75. WQBC 1.0 Quick Reference

| Instruction | Function              |
| ----------- | --------------------- |
| `PRINT`     | Output                |
| `SET`       | Variable assignment   |
| `LOAD`      | Variable loading      |
| `ADD`       | Addition              |
| `SUB`       | Subtraction           |
| `MUL`       | Multiplication        |
| `DIV`       | Division              |
| `IF`        | Conditional execution |
| `ENDIF`     | End conditional       |
| `EXIT`      | Stop execution        |

WQBC instructions are implementation-level operations and may evolve independently from source-level syntax.

---

## 76. WQA Package Quick Reference

A `.wqa` package contains a binary package structure consisting of:

```text
WQA Header
Manifest
Application / WQBC
Resources
```

The header contains package metadata such as:

```text
Magic
Version
Flags
Offsets
Sizes
Signature
CRC32
```

Current package version:

```text
1.0
```

---

## 77. WQA Development Workflow

A typical WQA development workflow is:

```text
1. Create a WQA project
2. Write .wq source code
3. Validate the wqa header
4. Compile the source
5. Generate WQBC bytecode
6. Create the .wqa package
7. Run the package
8. Test the application
9. Distribute the .wqa package
```

Example:

```text
Project/
├── wqa.json
└── app/
    └── main.wq
```

Build:

```text
wqa build .
```

Run:

```text
wqa run .\Project.wqa
```

---

## 78. WQA Ecosystem

WQA is part of the broader WinDroid ecosystem.

The ecosystem is designed around several components:

```text
WQA
 |
 +-- WQA Language
 |
 +-- WQA Compiler
 |
 +-- WQBC
 |
 +-- WQA Runtime
 |
 +-- .wqa Packages
 |
 +-- WSh
```

WQA provides the application and package technology.

WSh provides a shell environment and application management layer.

Together they form the foundation for WinDroid application development.

---

## 79. Implementation Status

The WQA 1.0 language specification is an early development specification.

Implemented language areas include the core syntax for:

* WQA source headers
* Variables
* Strings
* Integers
* Output
* Arithmetic
* Conditional execution
* Comparisons
* `else`
* WQBC compilation
* WQA application packaging

Other features described as future functionality should not be considered part of the current stable language.

---

## 80. Specification Conformance

A WQA implementation may be considered compatible with this specification when it correctly handles the required WQA 1.0 syntax and package conventions defined here.

A compatible implementation should:

1. Recognize the lowercase `wqa` source header.
2. Support `wet` variable assignment.
3. Support `print`.
4. Support `add`.
5. Support `sub`.
6. Support `mul`.
7. Support `div`.
8. Support `wif`.
9. Support `else`.
10. Support `endwif`.
11. Support the defined comparison operators.
12. Correctly process WQBC generated from valid WQA source.
13. Validate WQA package metadata.
14. Respect the declared WQA package version.

---

## 81. Specification Status

**WQA Language 1.0 is an early development specification.**

The language and bytecode format may change before the first stable release.

Applications should declare their required WQA version.

Developers should avoid undocumented functionality.

This specification describes the intended WQA 1.0 language and runtime architecture and may be updated as the implementation develops.

---

# Win Studio

WQA and WSh are projects of **Win Studio**.

WQA is intended to provide a simple application language, bytecode format, runtime, and package system for the WinDroid ecosystem.

**WQA Language Specification 1.0**

**Runtime:** WQBC
**Architectures:** x86_64, arm64
**Project:** WQA / WSh
**Publisher:** Win Studio
