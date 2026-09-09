package runtime

import (
	"strings"
	"testing"
)

func TestRuntimeError(t *testing.T) {
	err := NewRuntimeError(
		ErrInvalidBytecode,
		"test error",
		-1,
		10,
	)

	if err.Code != ErrInvalidBytecode {
		t.Fatalf(
			"expected code %s, got %s",
			ErrInvalidBytecode,
			err.Code,
		)
	}

	if err.Message != "test error" {
		t.Fatalf(
			"expected message %q, got %q",
			"test error",
			err.Message,
		)
	}

	if err.PC != 10 {
		t.Fatalf(
			"expected pc 10, got %d",
			err.PC,
		)
	}

	if err.Instruction != -1 {
		t.Fatalf(
			"expected instruction -1, got %d",
			err.Instruction,
		)
	}
}

func TestRuntimeErrorError(t *testing.T) {
	err := NewRuntimeError(
		ErrInvalidBytecode,
		"invalid program",
		3,
		12,
	)

	message := err.Error()

	if !strings.Contains(message, "WQBC-007") {
		t.Fatalf(
			"error message does not contain error code: %s",
			message,
		)
	}

	if !strings.Contains(message, "invalid program") {
		t.Fatalf(
			"error message does not contain error text: %s",
			message,
		)
	}

	if !strings.Contains(message, "instruction 3") {
		t.Fatalf(
			"error message does not contain instruction: %s",
			message,
		)
	}

	if !strings.Contains(message, "pc 12") {
		t.Fatalf(
			"error message does not contain pc: %s",
			message,
		)
	}
}

func TestUnknownOpcode(t *testing.T) {
	err := UnknownOpcode(
		0xFF,
		4,
		12,
	)

	if err.Code != ErrUnknownOpcode {
		t.Fatalf(
			"expected %s, got %s",
			ErrUnknownOpcode,
			err.Code,
		)
	}

	if !strings.Contains(
		err.Message,
		"unknown opcode 0xFF",
	) {
		t.Fatalf(
			"unexpected message: %s",
			err.Message,
		)
	}
}

func TestDivisionByZero(t *testing.T) {
	err := DivisionByZero(
		5,
		20,
	)

	if err.Code != ErrDivisionByZero {
		t.Fatalf(
			"expected %s, got %s",
			ErrDivisionByZero,
			err.Code,
		)
	}

	if err.Message != "division by zero" {
		t.Fatalf(
			"unexpected message: %s",
			err.Message,
		)
	}
}

func TestInvalidRegister(t *testing.T) {
	err := InvalidRegister(
		99,
		2,
		8,
	)

	if err.Code != ErrInvalidRegister {
		t.Fatalf(
			"expected %s, got %s",
			ErrInvalidRegister,
			err.Code,
		)
	}

	if !strings.Contains(
		err.Message,
		"invalid register R99",
	) {
		t.Fatalf(
			"unexpected message: %s",
			err.Message,
		)
	}
}

func TestInvalidConstant(t *testing.T) {
	err := InvalidConstant(
		42,
		7,
		15,
	)

	if err.Code != ErrInvalidConstant {
		t.Fatalf(
			"expected %s, got %s",
			ErrInvalidConstant,
			err.Code,
		)
	}

	if !strings.Contains(
		err.Message,
		"invalid constant 42",
	) {
		t.Fatalf(
			"unexpected message: %s",
			err.Message,
		)
	}
}

func TestInvalidOperand(t *testing.T) {
	err := InvalidOperand(
		"expected integer",
		3,
		9,
	)

	if err.Code != ErrInvalidOperand {
		t.Fatalf(
			"expected %s, got %s",
			ErrInvalidOperand,
			err.Code,
		)
	}

	if err.Message != "expected integer" {
		t.Fatalf(
			"unexpected message: %s",
			err.Message,
		)
	}
}

func TestInvalidInstruction(t *testing.T) {
	err := InvalidInstruction(
		"missing operand",
		2,
		6,
	)

	if err.Code != ErrInvalidInstruction {
		t.Fatalf(
			"expected %s, got %s",
			ErrInvalidInstruction,
			err.Code,
		)
	}

	if err.Message != "missing operand" {
		t.Fatalf(
			"unexpected message: %s",
			err.Message,
		)
	}
}

func TestInvalidBytecode(t *testing.T) {
	err := InvalidBytecode(
		"unexpected end of bytecode",
		25,
	)

	if err.Code != ErrInvalidBytecode {
		t.Fatalf(
			"expected %s, got %s",
			ErrInvalidBytecode,
			err.Code,
		)
	}

	if err.Instruction != -1 {
		t.Fatalf(
			"expected instruction -1, got %d",
			err.Instruction,
		)
	}

	if err.PC != 25 {
		t.Fatalf(
			"expected pc 25, got %d",
			err.PC,
		)
	}
}

func TestAllErrorCodesAreUnique(t *testing.T) {
	codes := []ErrorCode{
		ErrUnknownOpcode,
		ErrDivisionByZero,
		ErrInvalidRegister,
		ErrInvalidConstant,
		ErrInvalidOperand,
		ErrInvalidInstruction,
		ErrInvalidBytecode,
	}

	seen := make(map[ErrorCode]bool)

	for _, code := range codes {
		if seen[code] {
			t.Fatalf(
				"duplicate error code: %s",
				code,
			)
		}

		seen[code] = true
	}
}
