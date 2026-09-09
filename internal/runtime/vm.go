package runtime

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

type VM struct {
	Code []byte
	IP   int

	Variables map[string]interface{}
	Stack     []interface{}
}

func New(code []byte) *VM {
	return &VM{
		Code:      code,
		IP:        0,
		Variables: make(map[string]interface{}),
		Stack:     make([]interface{}, 0),
	}
}

func (vm *VM) readString() (string, error) {

	if vm.IP+2 > len(vm.Code) {
		return "", fmt.Errorf(
			"unexpected end of bytecode while reading string length",
		)
	}

	length := int(
		binary.BigEndian.Uint16(
			vm.Code[vm.IP : vm.IP+2],
		),
	)

	vm.IP += 2

	if vm.IP+length > len(vm.Code) {
		return "", fmt.Errorf(
			"unexpected end of bytecode while reading string",
		)
	}

	text := string(
		vm.Code[vm.IP : vm.IP+length],
	)

	vm.IP += length

	return text, nil
}

func (vm *VM) push(value interface{}) {
	vm.Stack = append(
		vm.Stack,
		value,
	)
}

func (vm *VM) pop() (interface{}, error) {

	if len(vm.Stack) == 0 {
		return nil, fmt.Errorf(
			"stack underflow",
		)
	}

	index := len(vm.Stack) - 1

	value := vm.Stack[index]

	vm.Stack = vm.Stack[:index]

	return value, nil
}

func (vm *VM) getNumber(value interface{}) (int, error) {

	switch v := value.(type) {

	case int:
		return v, nil

	case string:
		number, err := strconv.Atoi(v)

		if err != nil {
			return 0, fmt.Errorf(
				"%q is not a number",
				v,
			)
		}

		return number, nil

	default:
		return 0, fmt.Errorf(
			"value %v is not a number",
			value,
		)
	}
}

func (vm *VM) Run() {

	for vm.IP < len(vm.Code) {

		op := vm.Code[vm.IP]
		vm.IP++

		switch op {

		case OP_PUSH:

			value, err := vm.readString()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			if number, err := strconv.Atoi(value); err == nil {
				vm.push(number)
			} else {
				vm.push(value)
			}

		case OP_PRINT:

			value, err := vm.pop()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			fmt.Println(value)

		case OP_SET:

			name, err := vm.readString()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			value, err := vm.pop()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			vm.Variables[name] = value

		case OP_LOAD:

			name, err := vm.readString()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			value, exists := vm.Variables[name]

			if !exists {
				fmt.Printf(
					"Runtime error: undefined variable %q\n",
					name,
				)
				return
			}

			vm.push(value)

		case OP_ADD:

			vm.binaryOperation(
				func(a, b int) int {
					return a + b
				},
			)

		case OP_SUB:

			vm.binaryOperation(
				func(a, b int) int {
					return a - b
				},
			)

		case OP_MUL:

			vm.binaryOperation(
				func(a, b int) int {
					return a * b
				},
			)

		case OP_DIV:

			right, err := vm.pop()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			left, err := vm.pop()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			a, err := vm.getNumber(left)

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			b, err := vm.getNumber(right)

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			if b == 0 {
				fmt.Println(
					"Runtime error: division by zero",
				)
				return
			}

			vm.push(a / b)

		case OP_IF:

			operator, err := vm.readString()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			right, err := vm.pop()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			left, err := vm.pop()

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			l, err := vm.getNumber(left)

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			r, err := vm.getNumber(right)

			if err != nil {
				fmt.Println(
					"Runtime error:",
					err,
				)
				return
			}

			result := false

			switch operator {

			case ">":
				result = l > r

			case "<":
				result = l < r

			case "==":
				result = l == r

			default:
				fmt.Printf(
					"Runtime error: unknown comparison operator %q\n",
					operator,
				)
				return
			}

			if !result {
				vm.skipToElseOrEndIf()
			}

		case OP_ELSE:

			vm.skipToEndIf()

		case OP_ENDIF:

			continue

		case OP_LOOP:

			continue

		case OP_ENDLOOP:

			continue

		case OP_EXIT:

			return

		default:

			fmt.Printf(
				"Runtime error: unknown opcode 0x%02X\n",
				op,
			)

			return
		}
	}
}

func (vm *VM) skipToElseOrEndIf() {

	depth := 0

	for vm.IP < len(vm.Code) {

		op := vm.Code[vm.IP]
		vm.IP++

		switch op {

		case OP_IF:

			depth++

			if _, err := vm.readString(); err != nil {
				return
			}

		case OP_ELSE:

			if depth == 0 {
				return
			}

		case OP_ENDIF:

			if depth == 0 {
				return
			}

			depth--

		default:

			if !vm.skipInstructionOperands(op) {
				return
			}
		}
	}
}

func (vm *VM) skipToEndIf() {

	depth := 0

	for vm.IP < len(vm.Code) {

		op := vm.Code[vm.IP]
		vm.IP++

		switch op {

		case OP_IF:

			depth++

			if _, err := vm.readString(); err != nil {
				return
			}

		case OP_ENDIF:

			if depth == 0 {
				return
			}

			depth--

		default:

			if !vm.skipInstructionOperands(op) {
				return
			}
		}
	}
}

func (vm *VM) skipInstructionOperands(op byte) bool {

	switch op {

	case OP_PUSH,
		OP_SET,
		OP_LOAD,
		OP_IF:

		_, err := vm.readString()

		return err == nil

	default:
		return true
	}
}

func (vm *VM) binaryOperation(
	operation func(int, int) int,
) {

	right, err := vm.pop()

	if err != nil {
		fmt.Println(
			"Runtime error:",
			err,
		)
		return
	}

	left, err := vm.pop()

	if err != nil {
		fmt.Println(
			"Runtime error:",
			err,
		)
		return
	}

	a, err := vm.getNumber(left)

	if err != nil {
		fmt.Println(
			"Runtime error:",
			err,
		)
		return
	}

	b, err := vm.getNumber(right)

	if err != nil {
		fmt.Println(
			"Runtime error:",
			err,
		)
		return
	}

	vm.push(
		operation(a, b),
	)
}
