package runtime

import "fmt"

type ErrorCode string

const (
	ErrUnknownOpcode      ErrorCode = "WQBC-001"
	ErrDivisionByZero     ErrorCode = "WQBC-002"
	ErrInvalidRegister    ErrorCode = "WQBC-003"
	ErrInvalidConstant    ErrorCode = "WQBC-004"
	ErrInvalidOperand     ErrorCode = "WQBC-005"
	ErrInvalidInstruction ErrorCode = "WQBC-006"
	ErrInvalidBytecode    ErrorCode = "WQBC-007"
)

type RuntimeError struct {
	Code        ErrorCode
	Message     string
	Instruction int
	PC          int
}

func (e *RuntimeError) Error() string {
	if e.Instruction >= 0 {
		return fmt.Sprintf(
			"%s: %s (instruction %d, pc %d)",
			e.Code,
			e.Message,
			e.Instruction,
			e.PC,
		)
	}

	return fmt.Sprintf(
		"%s: %s (pc %d)",
		e.Code,
		e.Message,
		e.PC,
	)
}

func NewRuntimeError(
	code ErrorCode,
	message string,
	instruction int,
	pc int,
) *RuntimeError {
	return &RuntimeError{
		Code:        code,
		Message:     message,
		Instruction: instruction,
		PC:          pc,
	}
}

func UnknownOpcode(
	opcode byte,
	instruction int,
	pc int,
) *RuntimeError {
	return NewRuntimeError(
		ErrUnknownOpcode,
		fmt.Sprintf(
			"unknown opcode 0x%02X",
			opcode,
		),
		instruction,
		pc,
	)
}

func DivisionByZero(
	instruction int,
	pc int,
) *RuntimeError {
	return NewRuntimeError(
		ErrDivisionByZero,
		"division by zero",
		instruction,
		pc,
	)
}

func InvalidRegister(
	register int,
	instruction int,
	pc int,
) *RuntimeError {
	return NewRuntimeError(
		ErrInvalidRegister,
		fmt.Sprintf(
			"invalid register R%d",
			register,
		),
		instruction,
		pc,
	)
}

func InvalidConstant(
	constant int,
	instruction int,
	pc int,
) *RuntimeError {
	return NewRuntimeError(
		ErrInvalidConstant,
		fmt.Sprintf(
			"invalid constant %d",
			constant,
		),
		instruction,
		pc,
	)
}

func InvalidOperand(
	message string,
	instruction int,
	pc int,
) *RuntimeError {
	return NewRuntimeError(
		ErrInvalidOperand,
		message,
		instruction,
		pc,
	)
}

func InvalidInstruction(
	message string,
	instruction int,
	pc int,
) *RuntimeError {
	return NewRuntimeError(
		ErrInvalidInstruction,
		message,
		instruction,
		pc,
	)
}

func InvalidBytecode(
	message string,
	pc int,
) *RuntimeError {
	return NewRuntimeError(
		ErrInvalidBytecode,
		message,
		-1,
		pc,
	)
}
