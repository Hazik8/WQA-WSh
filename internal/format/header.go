package format

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
)

type Header struct {
	Magic           [4]byte
	MajorVersion    uint8
	MinorVersion    uint8
	Flags           uint16
	HeaderSize      uint32
	ManifestOffset  uint32
	ManifestSize    uint32
	AppOffset       uint64
	AppSize         uint64
	ResourcesOffset uint64
	ResourcesSize   uint64
	SignatureOffset uint32
	SignatureSize   uint32
	CRC32           uint32
}

func NewHeader() Header {
	var magic [4]byte
	copy(magic[:], []byte(Magic))

	return Header{
		Magic:        magic,
		MajorVersion: MajorVersion,
		MinorVersion: MinorVersion,
		HeaderSize:   HeaderSize,
	}
}

func (h *Header) MarshalBinary() ([]byte, error) {

	var buf bytes.Buffer

	err := binary.Write(
		&buf,
		binary.LittleEndian,
		h,
	)

	if err != nil {
		return nil, err
	}

	data := buf.Bytes()

	if len(data) != 64 {
		return nil, fmt.Errorf("wrong header size")
	}

	crc := crc32.ChecksumIEEE(data[:60])

	binary.LittleEndian.PutUint32(
		data[60:64],
		crc,
	)

	return data, nil
}
func ParseHeader(data []byte) (*Header, error) {

	if len(data) < int(HeaderSize) {
		return nil, fmt.Errorf("file is too small")
	}

	var header Header

	err := binary.Read(
		bytes.NewReader(data[:HeaderSize]),
		binary.LittleEndian,
		&header,
	)

	if err != nil {
		return nil, err
	}

	if string(header.Magic[:]) != Magic {
		return nil, fmt.Errorf("not a WQA file")
	}

	if header.MajorVersion != MajorVersion {
		return nil, fmt.Errorf(
			"unsupported version %d.%d",
			header.MajorVersion,
			header.MinorVersion,
		)
	}

	expectedCRC := binary.LittleEndian.Uint32(
		data[60:64],
	)

	actualCRC := crc32.ChecksumIEEE(
		data[:60],
	)

	if expectedCRC != actualCRC {
		return nil, fmt.Errorf(
			"CRC32 mismatch",
		)
	}

	return &header, nil
}
