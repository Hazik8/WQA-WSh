# WQA Language Specification

**WQA Language Version:** 1.0
**Runtime:** WQBC
**Architecture:** x86_64, arm64
**Project:** WQA / WSh
**Publisher:** Win Studio

---

## 1. About WQA

WQA is the application language and runtime environment used by the WinDroid application ecosystem.

WQA applications are designed to run through the WQA Runtime and can be packaged into `.wqa` application packages.

The language is intentionally small and simple.

Example:

```wqa
print "Hello, WQA!"
```

---

## 2. File Types

WQA uses several file types.

| Extension | Purpose                 |
| --------- | ----------------------- |
| `.wqa`    | WQA application package |
| `.wqbc`   | WQBC bytecode           |
| `.wq`     | WQA source file         |

A `.wqa` package can contain the application manifest, bytecode and resources.

---

## 3. Hello World

The simplest WQA program:

```wqa
print "Hello, WQA!"
```

Output:

```text
Hello, WQA!
```

---

## 4. Output

Use `print` to display text:

```wqa
print "Hello"
print "Welcome to WQA"
```

Each `print` statement writes a new line.

---

## 5. Variables

Variables can store numbers and strings.

Example:

```wqa
set score 100
set name "Calculator"
```

A variable can be loaded with:

```wqa
load score
```

Output:

```text
100
```

---

## 6. Numbers

WQA supports integer values.

Example:

```wqa
set score 100
set lives 3
set version 1
```

---

## 7. Strings

Strings are text values.

Example:

```wqa
set name "WinDroid"
set message "Hello!"
```

---

## 8. Addition

The `add` instruction adds two numbers.

```wqa
set a 10
set b 20

add a b result

load result
```

Output:

```text
30
```

The result is stored in the variable specified by the third argument.

---

## 9. Subtraction

The `sub` instruction subtracts one number from another.

```wqa
set a 30
set b 10

sub a b result

load result
```

Output:

```text
20
```

---

## 10. Multiplication

The `mul` instruction multiplies two numbers.

```wqa
set a 5
set b 4

mul a b result

load result
```

Output:

```text
20
```

---

## 11. Division

The `div` instruction divides two numbers.

```wqa
set a 20
set b 5

div a b result

load result
```

Output:

```text
4
```

Division by zero is not performed.

---

## 12. Conditions

WQA supports conditional execution using `if`.

Basic syntax:

```wqa
if score > 50
    print "Winner"
endif
```

Supported comparison operators:

```text
>
<
==
```

### Greater than

```wqa
if score > 50
    print "Winner"
endif
```

### Less than

```wqa
if score < 50
    print "Try again"
endif
```

### Equal

```wqa
if score == 100
    print "Perfect"
endif
```

---

## 13. End of Condition

Every `if` block must be closed with:

```wqa
endif
```

Example:

```wqa
set score 100

if score > 50
    print "Winner"
endif
```

---

## 14. Exit

The `exit` instruction terminates the program.

```wqa
print "Goodbye"
exit
```

---

# 15. Complete Example

```wqa
set score 100

print "Game started"

if score > 50
    print "Winner"
endif

print "Game finished"

exit
```

Expected output:

```text
Game started
Winner
Game finished
```

---

# 16. WQBC

WQA source code is executed by the WQA Runtime through WQBC bytecode.

The runtime contains a virtual machine that processes WQBC instructions.

Current instruction set:

| Instruction | Description                 |
| ----------- | --------------------------- |
| `PRINT`     | Print text                  |
| `SET`       | Create or update a variable |
| `LOAD`      | Load and display a variable |
| `ADD`       | Addition                    |
| `SUB`       | Subtraction                 |
| `MUL`       | Multiplication              |
| `DIV`       | Division                    |
| `IF`        | Conditional execution       |
| `ENDIF`     | End conditional block       |
| `EXIT`      | Stop execution              |

---

# 17. Variables in WQBC

Variables are stored by name.

Examples:

```text
score
name
result
lives
```

Numeric values are stored as integers.

String values are stored as strings.

---

# 18. WQA Application Manifest

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
  "entry": "app/main.wqa",
  "architecture": [
    "x86_64",
    "arm64"
  ]
}
```

### Manifest fields

| Field          | Description                     |
| -------------- | ------------------------------- |
| `wqa`          | WQA format version              |
| `id`           | Unique application identifier   |
| `name`         | Application name                |
| `version`      | Application version             |
| `publisher`    | Application publisher           |
| `runtime`      | Runtime used by the application |
| `entry`        | Application entry point         |
| `architecture` | Supported CPU architectures     |

---

# 19. Application IDs

Every installed WQA application should have a unique ID.

Example:

```text
windroid.Calculator
```

The ID is used by the package manager.

Example:

```text
wsh> run windroid.Calculator
```

---

# 20. WQA Packages

A WQA application can be distributed as:

```text
Calculator.wqa
```

The package can contain:

```text
manifest
bytecode
resources
application files
```

The WQA package format contains a binary header with:

```text
Magic
Version
Flags
Manifest
Application
Resources
Signature
CRC32
```

---

# 21. Runtime

WQA programs are executed by the WQA Runtime.

The runtime provides:

* WQBC virtual machine
* Variables
* Arithmetic
* Conditions
* Output
* Program termination

The runtime is responsible for validating and executing WQBC instructions.

---

# 22. Architecture

WQA applications can declare supported architectures.

Currently supported architecture identifiers:

```text
x86_64
arm64
```

Architecture support is declared in the manifest.

---

# 23. Versioning

WQA uses semantic versions for applications.

Example:

```text
1.0.0
```

Format:

```text
MAJOR.MINOR.PATCH
```

The WQA format itself currently uses:

```text
WQA 1.0
```

---

# 24. Comments

Comments are planned for a future version of the language.

Current WQA 1.0 implementations should not depend on comment syntax.

---

# 25. Functions

User-defined functions are planned for a future version.

Example future syntax:

```wqa
function hello
    print "Hello!"
end
```

This syntax is **not part of WQA 1.0**.

---

# 26. Future Features

The following features may be added in future WQA versions:

* Functions
* Boolean values
* More comparison operators
* Logical operators
* Loops
* Arrays
* Better string operations
* File APIs
* System APIs
* GUI APIs
* Networking APIs
* Application permissions
* Cryptographic signatures
* Package resources
* Cross-platform runtime

Future features are not guaranteed to be compatible with WQA 1.0.

---

# 27. Design Philosophy

WQA is designed to be:

* Small
* Fast
* Easy to learn
* Easy to implement
* Suitable for application packaging
* Suitable for the WinDroid ecosystem

The language intentionally starts with a small instruction set.

More advanced functionality can be added through the runtime and future WQA versions.

---

# 28. Example Calculator

A minimal calculator-style program:

```wqa
set a 10
set b 5

add a b result

print "Result:"
load result

exit
```

Output:

```text
Result:
15
```

---

# 29. Compatibility

A WQA 1.0 runtime should support the instructions defined in this specification.

Programs should avoid relying on undocumented instructions.

Applications should declare their WQA version in the package manifest.

---

# 30. Status

**WQA Language 1.0 is an early development specification.**

The language and bytecode format may change before the first stable release.

WQA applications should therefore declare their required WQA version.

---

## Win Studio

WQA and WSh are projects of **Win Studio**.
