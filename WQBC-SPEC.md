# WQBC 1.0 Specification

**WQA Bytecode Specification**

Version: **1.0**
Status: **Draft for WQA 1.0**

---

## 1. Overview

WQBC is the bytecode format used by the WQA runtime.

WQBC provides a stable interface between:

WQA Compiler
     ↓
   WQBC
     ↓
WQA Runtime / VM

The compiler produces WQBC instructions.

The runtime executes WQBC instructions.

The runtime MUST NOT depend on the source `.wq` syntax.

---

## 2. File Types

WQA uses three primary file types.

### `.wq`

WQA source code.

Example:

set x = 10
set y = 20
add x y result
print result

The source file is processed by the lexer, parser and compiler.

---

### `.wqbc`

Compiled WQA bytecode.

A `.wqbc` file contains executable bytecode intended for the WQA runtime.

It does not contain the original source code unless explicitly embedded by a future debug feature.

---

### `.wqa`

WQA application package.

A `.wqa` package contains:

Manifest
Application data
Resources
Optional signature

The package may contain WQBC produced from `.wq` source.

---

## 3. Versioning

WQBC has its own version number.

WQBC versioning is independent from:

* WQA language version
* WQA package version
* WSh version

Example:

WQA language: 1.0
WQBC:         1.0
Package:      1.0.0
WSh:          0.4.0

A runtime MUST reject bytecode versions that it does not support.

---

## 4. Execution Model

WQBC uses a register-based execution model.

The VM contains a fixed set of virtual registers.

Registers are identified by numeric indexes.

Example:

R0
R1
R2
R3

Instructions read values from registers and may write results to registers.

---

## 5. Instruction Structure

Each instruction consists of:

opcode
operands

Conceptually:

[ OPCODE ][ OPERAND 1 ][ OPERAND 2 ] ...

The exact binary encoding is implementation-defined for WQBC 1.0 unless explicitly specified by the bytecode container format.

Each opcode has a unique numeric identifier.

---

# 6. Opcode Table

## 6.1 Control

| Opcode | Name   | Description           |
| ------ | ------ | --------------------- |
| `0x00` | `NOP`  | Performs no operation |
| `0x01` | `HALT` | Stops VM execution    |

---

## 6.2 Data

| Opcode | Name         | Description                          |
| ------ | ------------ | ------------------------------------ |
| `0x10` | `LOAD_CONST` | Loads a constant into a register     |
| `0x11` | `MOVE`       | Copies one register value to another |

---

## 6.3 Arithmetic

| Opcode | Name  | Description           |
| ------ | ----- | --------------------- |
| `0x20` | `ADD` | Adds two values       |
| `0x21` | `SUB` | Subtracts two values  |
| `0x22` | `MUL` | Multiplies two values |
| `0x23` | `DIV` | Divides two values    |

---

## 6.4 Output

| Opcode | Name    | Description                      |
| ------ | ------- | -------------------------------- |
| `0x30` | `PRINT` | Writes a value to runtime output |

---

# 7. NOP

Opcode:

0x00

Operation:

NOP

`NOP` performs no operation.

It does not modify registers or constants.

---

# 8. HALT

Opcode:

0x01

Operation:

HALT

Stops execution of the current WQBC program.

`HALT` does not represent a runtime error.

---

# 9. LOAD_CONST

Opcode:

0x10

Operation:

LOAD_CONST destination, constant

Example:

LOAD_CONST R0, 10

The value of the selected constant is loaded into the destination register.

---

# 10. MOVE

Opcode:

0x11

Operation:

MOVE destination, source

Example:

MOVE R1, R0

After execution:

R1 = R0

The source register is not modified.

---

# 11. ADD

Opcode:

0x20

Operation:

ADD destination, left, right

Example:

ADD R2, R0, R1

Conceptually:

R2 = R0 + R1

The source registers remain unchanged.

---

# 12. SUB

Opcode:

0x21

Operation:

SUB destination, left, right

Example:

SUB R2, R0, R1

Conceptually:

R2 = R0 - R1

---

# 13. MUL

Opcode:

0x22

Operation:

MUL destination, left, right

Example:

MUL R2, R0, R1

Conceptually:

R2 = R0 * R1

---

# 14. DIV

Opcode:

0x23

Operation:

DIV destination, left, right

Example:

DIV R2, R0, R1

Conceptually:

R2 = R0 / R1

Division by zero MUST produce a runtime error.

---

# 15. PRINT

Opcode:

0x30

Operation:

PRINT source

Example:

PRINT R0

The runtime writes the value stored in the selected register to standard output.

---

# 16. Example Program

WQA source:

set x = 10
set y = 20
add x y result
print result

Conceptual AST:

Program
├── SetStatement
│   ├── x
│   └── 10
│
├── SetStatement
│   ├── y
│   └── 20
│
├── BinaryStatement
│   ├── ADD
│   ├── x
│   ├── y
│   └── result
│
└── PrintStatement
    └── result

Conceptual WQBC:

LOAD_CONST R0, 10
LOAD_CONST R1, 20
ADD R2, R0, R1
PRINT R2
HALT

---

# 17. Constants

A WQBC program may contain a constant pool.

The constant pool stores values that can be referenced by instructions.

Initial supported constant types:

INTEGER
STRING

Example:

Constant 0 = 10
Constant 1 = 20
Constant 2 = "Hello"

`LOAD_CONST` references constants from this pool.

---

# 18. Runtime Errors

Runtime errors MUST be represented by a structured error system.

An error SHOULD contain:

Error code
Message
Instruction
Program counter

Example:

WQBC-001: division by zero

The runtime MUST NOT silently continue after a fatal bytecode execution error.

---

# 19. Invalid Opcodes

If the VM encounters an opcode that is not defined by the supported WQBC version, execution MUST stop.

Example:

WQBC-002: unknown opcode 0xFF

The runtime MUST NOT interpret an unknown opcode as another instruction.

---

# 20. Register Errors

Access to an invalid register MUST produce a runtime error.

Example:

WQBC-003: invalid register R99

The VM MUST NOT access memory outside the valid register range.

---

# 21. Type Errors

Operations MUST validate operand types.

For example, an arithmetic operation cannot operate on unsupported values.

Example:

WQBC-004: invalid operand type

The exact type system will be expanded in a future WQBC specification.

---

# 22. Program Termination

A valid program SHOULD end with:

HALT

If execution reaches the end of a valid bytecode stream without `HALT`, the runtime MAY terminate execution normally.

The exact container-level behavior will be defined by the WQBC binary format.

---

# 23. Compatibility

A WQBC runtime MUST identify the bytecode version before execution.

The runtime MUST reject unsupported major versions.

Minor versions MAY introduce backward-compatible features.

Example:

WQBC 1.0

A runtime supporting WQBC 1.x MAY support:

1.0
1.1
1.2

subject to compatibility rules.

A runtime supporting WQBC 1.x MUST NOT automatically execute:

2.0

unless explicit compatibility support exists.

---

# 24. Compiler Responsibilities

The compiler is responsible for:

1. Validating the AST.
2. Allocating registers.
3. Creating the constant pool.
4. Generating valid opcodes.
5. Producing valid WQBC.
6. Ensuring all referenced registers and constants exist.

The compiler MUST NOT generate undefined opcodes.

---

# 25. Runtime Responsibilities

The runtime is responsible for:

1. Validating WQBC.
2. Checking opcode validity.
3. Checking operand validity.
4. Executing instructions.
5. Maintaining register state.
6. Handling runtime errors.
7. Stopping execution when required.

The runtime MUST NOT parse `.wq` source code.

---

# 26. Validation

Before execution, the runtime SHOULD validate:

WQBC version
Opcode validity
Operand count
Register indexes
Constant indexes
Program boundaries

Invalid bytecode MUST NOT be executed.

---

# 27. Security

WQBC MUST be treated as untrusted input.

A runtime MUST NOT assume that bytecode was generated by a trusted compiler.

The runtime MUST validate all indexes and sizes before accessing data.

Malformed bytecode MUST result in an error rather than an out-of-bounds access or crash.

---

# 28. Reserved Opcode Space

The following ranges are reserved for future use:

0x02 - 0x0F
0x12 - 0x1F
0x24 - 0x2F
0x31 - 0xFF

Implementations MUST NOT assign private instructions to these ranges while claiming WQBC 1.0 compatibility.

Future specifications may define additional instructions.

---

# 29. Future Instructions

Possible future WQBC instructions include:

JUMP
JUMP_IF
CALL
RETURN
COMPARE
AND
OR
NOT
LOAD_GLOBAL
STORE_GLOBAL

These are NOT part of WQBC 1.0.

A WQBC 1.0 runtime MUST NOT assume that they exist.

---

# 30. WQBC 1.0 Minimal Instruction Set

The complete mandatory instruction set for WQBC 1.0 is:

0x00 NOP
0x01 HALT

0x10 LOAD_CONST
0x11 MOVE

0x20 ADD
0x21 SUB
0x22 MUL
0x23 DIV

0x30 PRINT

This is the minimum instruction set required for a WQBC 1.0 implementation.

---

# 31. Specification Status

WQBC 1.0 is the bytecode specification targeted by WQA 1.0.

Changes to the instruction set or execution model MUST be documented and versioned.

Experimental instructions MUST NOT be presented as standard WQBC 1.0 instructions.
