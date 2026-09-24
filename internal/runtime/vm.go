package runtime

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type RepeatFrame struct {
	Remaining int
	BodyStart int
	Exit      int
}

type VM struct {
	Code        []byte
	IP          int
	Variables   map[string]interface{}
	Stack       []interface{}
	RepeatStack []RepeatFrame
	CallStack   []int
}

func NewVM(code []byte) *VM {
	return &VM{
		Code:        code,
		Variables:   make(map[string]interface{}),
		Stack:       make([]interface{}, 0),
		RepeatStack: make([]RepeatFrame, 0),
		CallStack:   make([]int, 0),
	}
}

func (vm *VM) readString() string {
	if vm.IP+2 > len(vm.Code) {
		panic("invalid bytecode: missing string length")
	}

	length := int(binary.BigEndian.Uint16(vm.Code[vm.IP:]))
	vm.IP += 2

	if vm.IP+length > len(vm.Code) {
		panic("invalid bytecode: string exceeds code size")
	}

	value := string(vm.Code[vm.IP : vm.IP+length])
	vm.IP += length

	return value
}

func (vm *VM) push(value interface{}) {
	vm.Stack = append(vm.Stack, value)
}

func (vm *VM) pop() (interface{}, error) {
	if len(vm.Stack) == 0 {
		return nil, fmt.Errorf("stack underflow")
	}

	last := len(vm.Stack) - 1
	value := vm.Stack[last]
	vm.Stack = vm.Stack[:last]

	return value, nil
}

func getNumber(value interface{}) (float64, error) {
	switch v := value.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		n, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return 0, fmt.Errorf("not a number: %v", value)
		}
		return n, nil
	default:
		return 0, fmt.Errorf("not a number: %v", value)
	}
}

func normalizeNumber(value float64) interface{} {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return value
	}

	if value == math.Trunc(value) &&
		value >= float64(int(^uint(0)>>1)*-1) &&
		value <= float64(int(^uint(0)>>1)) {
		return int(value)
	}

	return value
}

func (vm *VM) Run() {
	for vm.IP < len(vm.Code) {
		op := vm.Code[vm.IP]
		vm.IP++

		switch op {
		case OP_PUSH:
			value := vm.readString()

			if number, err := strconv.ParseFloat(value, 64); err == nil {
				vm.push(normalizeNumber(number))
			} else {
				vm.push(value)
			}

		case OP_LOAD:
			name := vm.readString()

			value, ok := vm.Variables[name]
			if !ok {
				fmt.Printf("undefined variable: %s\n", name)
				return
			}

			vm.push(value)

		case OP_SET:
			name := vm.readString()

			value, err := vm.pop()
			if err != nil {
				fmt.Println(err)
				return
			}

			vm.Variables[name] = value

		case OP_ADD:
			if !vm.binaryOperation(func(a, b float64) float64 {
				return a + b
			}) {
				return
			}

		case OP_SUB:
			if !vm.binaryOperation(func(a, b float64) float64 {
				return a - b
			}) {
				return
			}

		case OP_MUL:
			if !vm.binaryOperation(func(a, b float64) float64 {
				return a * b
			}) {
				return
			}

		case OP_DIV:
			if len(vm.Stack) < 2 {
				fmt.Println("stack underflow")
				return
			}

			rightValue, _ := vm.pop()
			leftValue, _ := vm.pop()

			right, err := getNumber(rightValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			left, err := getNumber(leftValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			if right == 0 {
				fmt.Println("division by zero")
				return
			}

			vm.push(normalizeNumber(left / right))

		case OP_IDIV:
			if len(vm.Stack) < 2 {
				fmt.Println("stack underflow")
				return
			}

			rightValue, _ := vm.pop()
			leftValue, _ := vm.pop()

			right, err := getNumber(rightValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			left, err := getNumber(leftValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			if right == 0 {
				fmt.Println("division by zero")
				return
			}

			vm.push(normalizeNumber(math.Floor(left / right)))

		case OP_MOD:
			if len(vm.Stack) < 2 {
				fmt.Println("stack underflow")
				return
			}

			rightValue, _ := vm.pop()
			leftValue, _ := vm.pop()

			right, err := getNumber(rightValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			left, err := getNumber(leftValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			if right == 0 {
				fmt.Println("division by zero")
				return
			}

			result := left - math.Floor(left/right)*right
			vm.push(normalizeNumber(result))

		case OP_NEG:
			value, err := vm.pop()
			if err != nil {
				fmt.Println(err)
				return
			}

			number, err := getNumber(value)
			if err != nil {
				fmt.Println(err)
				return
			}

			vm.push(normalizeNumber(-number))

		case OP_PRINT:
			value, err := vm.pop()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(value)

		case OP_TIME:
			now := time.Now()
			fmt.Println(now.Format("15:04:05"))

		case OP_DATE:
			now := time.Now()
			fmt.Println(now.Format("02.01.2006"))

		case OP_INPUT:
			name := vm.readString()

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("> ")

			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			vm.Variables[name] = strings.TrimSpace(input)

		case OP_CLEAR:
			fmt.Print("\033[H\033[2J")

		case OP_WAIT:
			value, err := vm.pop()
			if err != nil {
				fmt.Println(err)
				return
			}

			seconds, err := getNumber(value)
			if err != nil {
				fmt.Println(err)
				return
			}

			if seconds < 0 {
				fmt.Println("wait duration cannot be negative")
				return
			}

			time.Sleep(time.Duration(seconds * float64(time.Second)))

		case OP_IF:
			operator := vm.readString()

			rightValue, err := vm.pop()
			if err != nil {
				fmt.Println(err)
				return
			}

			leftValue, err := vm.pop()
			if err != nil {
				fmt.Println(err)
				return
			}

			left, err := getNumber(leftValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			right, err := getNumber(rightValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			result := compare(left, right, operator)

			if !result {
				vm.skipToElseOrEndIf()
			}

		case OP_ELSE:
			vm.skipToEndIf()

		case OP_ENDIF:

		case OP_LOOP:
			operator := vm.readString()
			exit := int(binary.BigEndian.Uint32(vm.Code[vm.IP:]))
			vm.IP += 4

			if len(vm.Stack) < 2 {
				fmt.Println("stack underflow")
				return
			}

			rightValue, _ := vm.pop()
			leftValue, _ := vm.pop()

			left, err := getNumber(leftValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			right, err := getNumber(rightValue)
			if err != nil {
				fmt.Println(err)
				return
			}

			if !compare(left, right, operator) {
				vm.IP = exit
			}

		case OP_JUMP:
			vm.IP = int(binary.BigEndian.Uint32(vm.Code[vm.IP:]))

		case OP_REPEAT:
			exit := int(binary.BigEndian.Uint32(vm.Code[vm.IP:]))
			vm.IP += 4

			value, err := vm.pop()
			if err != nil {
				fmt.Println(err)
				return
			}

			count, err := getNumber(value)
			if err != nil {
				fmt.Println(err)
				return
			}

			if count != math.Trunc(count) {
				fmt.Println("repeat count must be an integer")
				return
			}

			n := int(count)

			if n <= 0 {
				vm.IP = exit
				continue
			}

			vm.RepeatStack = append(vm.RepeatStack, RepeatFrame{
				Remaining: n,
				BodyStart: vm.IP,
				Exit:      exit,
			})

		case OP_ENDREPEAT:
			if len(vm.RepeatStack) == 0 {
				fmt.Println("repeat stack underflow")
				return
			}

			index := len(vm.RepeatStack) - 1
			frame := &vm.RepeatStack[index]

			frame.Remaining--

			if frame.Remaining > 0 {
				vm.IP = frame.BodyStart
			} else {
				vm.RepeatStack = vm.RepeatStack[:index]
			}

		case OP_CALL:
			address := int(binary.BigEndian.Uint32(vm.Code[vm.IP:]))
			vm.IP += 4

			vm.CallStack = append(vm.CallStack, vm.IP)
			vm.IP = address

		case OP_RET:
			if len(vm.CallStack) == 0 {
				return
			}

			index := len(vm.CallStack) - 1
			vm.IP = vm.CallStack[index]
			vm.CallStack = vm.CallStack[:index]

		case OP_ENDLOOP:

		case OP_EXIT:
			return

		default:
			fmt.Printf("unknown opcode: 0x%02X\n", op)
			return
		}
	}
}

func compare(left, right float64, operator string) bool {
	switch operator {
	case ">":
		return left > right
	case "<":
		return left < right
	case "==":
		return left == right
	case "!=":
		return left != right
	case ">=":
		return left >= right
	case "<=":
		return left <= right
	default:
		return false
	}
}

func (vm *VM) binaryOperation(operation func(float64, float64) float64) bool {
	rightValue, err := vm.pop()
	if err != nil {
		fmt.Println(err)
		return false
	}

	leftValue, err := vm.pop()
	if err != nil {
		fmt.Println(err)
		return false
	}

	right, err := getNumber(rightValue)
	if err != nil {
		fmt.Println(err)
		return false
	}

	left, err := getNumber(leftValue)
	if err != nil {
		fmt.Println(err)
		return false
	}

	vm.push(normalizeNumber(operation(left, right)))
	return true
}

func (vm *VM) skipToElseOrEndIf() {
	depth := 0

	for vm.IP < len(vm.Code) {
		op := vm.Code[vm.IP]
		vm.IP++

		switch op {
		case OP_IF:
			depth++
			vm.skipInstructionOperands(op)

		case OP_ENDIF:
			if depth == 0 {
				return
			}
			depth--

		case OP_ELSE:
			if depth == 0 {
				return
			}

		default:
			vm.skipInstructionOperands(op)
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
			vm.skipInstructionOperands(op)

		case OP_ENDIF:
			if depth == 0 {
				return
			}
			depth--

		default:
			vm.skipInstructionOperands(op)
		}
	}
}

func (vm *VM) skipInstructionOperands(op byte) {
	switch op {
	case OP_PUSH,
		OP_SET,
		OP_LOAD,
		OP_IF,
		OP_INPUT:

		if vm.IP+2 > len(vm.Code) {
			vm.IP = len(vm.Code)
			return
		}

		length := int(binary.BigEndian.Uint16(vm.Code[vm.IP:]))
		vm.IP += 2 + length

	case OP_JUMP,
		OP_CALL,
		OP_REPEAT:
		vm.IP += 4

	case OP_LOOP:
		if vm.IP+2 > len(vm.Code) {
			vm.IP = len(vm.Code)
			return
		}

		length := int(binary.BigEndian.Uint16(vm.Code[vm.IP:]))
		vm.IP += 2 + length
		vm.IP += 4
	}
}

func New(code []byte) *VM {
	return NewVM(code)
}
