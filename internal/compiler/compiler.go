package compiler

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"windroid/wqa/internal/runtime"
)

func writeString(out *bytes.Buffer, text string) {

	out.WriteByte(
		byte(len(text)),
	)

	out.WriteString(text)
}

func Compile(path string) ([]byte, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var output bytes.Buffer

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := strings.TrimSpace(
			scanner.Text(),
		)

		// PRINT
		if strings.HasPrefix(line, "print ") {

			text := strings.TrimPrefix(
				line,
				"print ",
			)

			output.WriteByte(
				runtime.OP_PRINT,
			)

			writeString(
				&output,
				text,
			)
		}

		// SET
		if strings.HasPrefix(line, "set ") {

			data := strings.TrimPrefix(
				line,
				"set ",
			)

			parts := strings.SplitN(
				data,
				" = ",
				2,
			)

			if len(parts) == 2 {

				name := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				output.WriteByte(
					runtime.OP_SET,
				)

				writeString(
					&output,
					name,
				)

				writeString(
					&output,
					value,
				)
			}
		}

		// LOAD
		if strings.HasPrefix(line, "load ") {

			name := strings.TrimSpace(
				strings.TrimPrefix(
					line,
					"load ",
				),
			)

			output.WriteByte(
				runtime.OP_LOAD,
			)

			writeString(
				&output,
				name,
			)
		}

		// ADD / SUB / MUL / DIV

		operations := map[string]byte{
			"add ": runtime.OP_ADD,
			"sub ": runtime.OP_SUB,
			"mul ": runtime.OP_MUL,
			"div ": runtime.OP_DIV,
		}

		for command, opcode := range operations {

			if strings.HasPrefix(line, command) {

				args := strings.Fields(
					line[len(command):],
				)

				if len(args) == 3 {

					output.WriteByte(
						opcode,
					)

					writeString(
						&output,
						args[0],
					)

					writeString(
						&output,
						args[1],
					)

					writeString(
						&output,
						args[2],
					)
				}
			}
			if strings.HasPrefix(line, "if ") {

				data := strings.TrimPrefix(
					line,
					"if ",
				)

				args := strings.Fields(data)

				if len(args) == 3 {

					output.WriteByte(
						runtime.OP_IF,
					)

					writeString(&output, args[0])
					writeString(&output, args[1])
					writeString(&output, args[2])
				}
			}

			if line == "endif" {

				output.WriteByte(
					runtime.OP_ENDIF,
				)
			}
			if line == "else" {

				output.WriteByte(
					runtime.OP_ELSE,
				)
			}

			if strings.HasPrefix(line, "loop ") {

				name := strings.TrimSpace(
					strings.TrimPrefix(line, "loop "),
				)

				output.WriteByte(
					runtime.OP_LOOP,
				)

				writeString(
					&output,
					name,
				)
			}

			if line == "endloop" {

				output.WriteByte(
					runtime.OP_ENDLOOP,
				)
			}
		}

		// EXIT
		if line == "exit" {

			output.WriteByte(
				runtime.OP_EXIT,
			)
		}
	}

	return output.Bytes(), scanner.Err()
}
