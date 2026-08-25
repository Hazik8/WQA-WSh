package format

const (
	Magic = "WQA\x00"

	MajorVersion byte = 1
	MinorVersion byte = 0

	HeaderSize uint32 = 64

	FlagCompressed uint16 = 1 << 0
	FlagSigned     uint16 = 1 << 1
	FlagEncrypted  uint16 = 1 << 2
	FlagExecutable uint16 = 1 << 3
)
