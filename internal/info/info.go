package info

import (
	"fmt"
	"os"

	"windroid/wqa/internal/format"
)

func Show(path string) error {

	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	header, err := format.ParseHeader(data)

	if err != nil {
		return err
	}

	fmt.Println("WQA Package Information")
	fmt.Println("-----------------------")
	fmt.Println("Magic: WQA")
	fmt.Printf("Version: %d.%d\n",
		header.MajorVersion,
		header.MinorVersion,
	)

	fmt.Printf("Header size: %d bytes\n",
		header.HeaderSize,
	)

	fmt.Printf("Flags: %d\n",
		header.Flags,
	)

	fmt.Printf("App offset: %d\n",
		header.AppOffset,
	)

	fmt.Printf("App size: %d bytes\n",
		header.AppSize,
	)

	fmt.Println("Status: OK")

	return nil
}
