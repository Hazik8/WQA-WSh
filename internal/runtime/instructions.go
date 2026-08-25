package runtime

const (
	OP_PRINT byte = 0x01
	OP_EXIT  byte = 0x02

	OP_SET  byte = 0x03
	OP_LOAD byte = 0x04

	OP_ADD byte = 0x05
	OP_SUB byte = 0x06
	OP_MUL byte = 0x07
	OP_DIV byte = 0x08

	OP_IF    byte = 0x09
	OP_ENDIF byte = 0x0A
	OP_ELSE  byte = 0x0B

	OP_LOOP    byte = 0x0C
	OP_ENDLOOP byte = 0x0D
)
