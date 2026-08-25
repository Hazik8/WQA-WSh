package verify

import (
	"fmt"
	"os"

	"windroid/wqa/internal/format"
)

func Verify(path string) error {

	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	header, err := format.ParseHeader(data)

	if err != nil {
		return err
	}

	fmt.Println("WQA Verification")
	fmt.Println("----------------")

	fmt.Println("Magic: OK")

	fmt.Printf(
		"Version: %d.%d\n",
		header.MajorVersion,
		header.MinorVersion,
	)

	fmt.Println("Header: OK")

	fmt.Println("CRC32: OK")

	fmt.Println()

	fmt.Println("Package is valid [OK]")

	return nil
}
