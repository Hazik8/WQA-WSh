package main

import "windroid/wqa/internal/runtime"

func main() {

	code := []byte{
		0x01,
		0x05,
		'H', 'e', 'l', 'l', 'o',
		0x02,
	}

	vm := runtime.New(code)

	vm.Run()
}
