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

	OP_LOOP      byte = 0x0C
	OP_ENDLOOP   byte = 0x0D
	OP_PUSH      byte = 0x0E
	OP_JUMP      byte = 0x0F
	OP_REPEAT    byte = 0x10
	OP_ENDREPEAT byte = 0x11
	OP_CALL      byte = 0x12
	OP_RET       byte = 0x13

	OP_TIME byte = 0x17
	OP_DATE byte = 0x18

	OP_INPUT byte = 0x14
	OP_CLEAR byte = 0x15
	OP_WAIT  byte = 0x16
)
