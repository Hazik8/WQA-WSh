package logger

import "fmt"

func Info(message string) {
	fmt.Println("[INFO]", message)
}

func Success(message string) {
	fmt.Println("[OK]", message)
}

func Error(message string) {
	fmt.Println("[ERROR]", message)
}

func Warning(message string) {
	fmt.Println("[WARN]", message)
}
