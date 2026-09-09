package compiler

type Emitter struct {
	code []byte
}

func NewEmitter() *Emitter {
	return &Emitter{
		code: make([]byte, 0),
	}
}

func (e *Emitter) Emit(op byte) {
	e.code = append(
		e.code,
		op,
	)
}

func (e *Emitter) EmitBytes(data []byte) {
	e.code = append(
		e.code,
		data...,
	)
}

func (e *Emitter) EmitString(value string) {

	length := uint16(len(value))

	// WQBC 1.0:
	// строка = uint16 длина + UTF-8 данные.
	e.code = append(
		e.code,
		byte(length>>8),
		byte(length),
	)

	e.code = append(
		e.code,
		[]byte(value)...,
	)
}

func (e *Emitter) Bytes() []byte {
	return e.code
}
