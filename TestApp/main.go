package main

import (
	"syscall"
	"unsafe"
)

var (
	user32     = syscall.NewLazyDLL("user32.dll")
	messageBox = user32.NewProc("MessageBoxW")
)

func main() {
	title, _ := syscall.UTF16PtrFromString("WinDroid TestApp")
	text, _ := syscall.UTF16PtrFromString(
		"TestApp успешно запущен через WSh!",
	)

	messageBox.Call(
		0,
		uintptr(unsafe.Pointer(text)),
		uintptr(unsafe.Pointer(title)),
		0,
	)
}
