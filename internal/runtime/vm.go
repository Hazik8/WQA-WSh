package runtime

import (
	"fmt"
	"strconv"
)

type VM struct {
	Code []byte
	IP   int

	Variables map[string]interface{}
}

func New(code []byte) *VM {

	return &VM{
		Code:      code,
		IP:        0,
		Variables: make(map[string]interface{}),
	}
}

func (vm *VM) readString() string {

	length := int(vm.Code[vm.IP])

	vm.IP++

	text := string(
		vm.Code[vm.IP : vm.IP+length],
	)

	vm.IP += length

	return text
}

func (vm *VM) getNumber(name string) int {

	value, exists := vm.Variables[name]

	if exists {

		switch v := value.(type) {

		case int:
			return v

		case string:
			number, _ := strconv.Atoi(v)
			return number
		}
	}

	number, _ := strconv.Atoi(name)

	return number
}

func (vm *VM) Run() {

	for vm.IP < len(vm.Code) {

		op := vm.Code[vm.IP]

		vm.IP++

		switch op {

		case OP_PRINT:

			text := vm.readString()

			fmt.Println(text)

		case OP_SET:

			name := vm.readString()

			value := vm.readString()

			if number, err := strconv.Atoi(value); err == nil {

				vm.Variables[name] = number

			} else {

				vm.Variables[name] = value
			}

		case OP_LOAD:

			name := vm.readString()

			value, exists := vm.Variables[name]

			if exists {
				fmt.Println(value)
			} else {
				fmt.Println("")
			}

		case OP_ADD:

			a := vm.readString()
			b := vm.readString()
			result := vm.readString()

			vm.Variables[result] =
				vm.getNumber(a) +
					vm.getNumber(b)

		case OP_SUB:

			a := vm.readString()
			b := vm.readString()
			result := vm.readString()

			vm.Variables[result] =
				vm.getNumber(a) -
					vm.getNumber(b)

		case OP_MUL:

			a := vm.readString()
			b := vm.readString()
			result := vm.readString()

			vm.Variables[result] =
				vm.getNumber(a) *
					vm.getNumber(b)

		case OP_DIV:

			a := vm.readString()
			b := vm.readString()
			result := vm.readString()

			divisor := vm.getNumber(b)

			if divisor != 0 {

				vm.Variables[result] =
					vm.getNumber(a) /
						divisor
			}

		case OP_IF:

			left := vm.readString()

			operator := vm.readString()

			right := vm.readString()

			l := vm.getNumber(left)
			r := vm.getNumber(right)

			result := false

			switch operator {

			case ">":
				result = l > r

			case "<":
				result = l < r

			case "==":
				result = l == r
			}

			if !result {

				for vm.IP < len(vm.Code) {

					if vm.Code[vm.IP] == OP_ENDIF {

						vm.IP++
						break
					}

					vm.IP++
				}
			}

		case OP_ENDIF:

			// конец условия
			continue

		case OP_EXIT:

			return

		default:

			fmt.Println(
				"Unknown opcode:",
				op,
			)

			return
		}
	}
}
