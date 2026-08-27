package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"windroid/wqa/internal/builder"
	"windroid/wqa/internal/info"
	"windroid/wqa/internal/installer"
	"windroid/wqa/internal/list"
	"windroid/wqa/internal/remove"
	"windroid/wqa/internal/run"
	"windroid/wqa/internal/setup"
	"windroid/wqa/internal/updater"
	"windroid/wqa/internal/verify"
)

const version = "1.0.0"

type Manifest struct {
	WQA          string   `json:"wqa"`
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Publisher    string   `json:"publisher"`
	Runtime      string   `json:"runtime"`
	Entry        string   `json:"entry"`
	Architecture []string `json:"architecture"`
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	switch os.Args[1] {

	case "--version":
		fmt.Println("WQA CLI", version)

	case "init":
		if len(os.Args) < 3 {
			fmt.Println("Использование: wqa init <name>")
			return
		}

		createProject(os.Args[2])

	case "--help", "-h", "help", "?":
		help()

	case "build":

		project := "."

		if len(os.Args) >= 3 {
			project = os.Args[2]
		}

		project, err := filepath.Abs(project)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		err = builder.Build(project)

		if err != nil {
			fmt.Println("Ошибка:", err)
		}

	case "info":

		if len(os.Args) < 3 {
			fmt.Println("Использование: wqa info <file>")
			return
		}

		err := info.Show(os.Args[2])

		if err != nil {
			fmt.Println("Ошибка:", err)
		}

	case "verify":

		if len(os.Args) < 3 {
			fmt.Println(
				"Использование: wqa verify <file>",
			)
			return
		}

		err := verify.Verify(os.Args[2])

		if err != nil {
			fmt.Println(
				"Ошибка:",
				err,
			)
		}

	case "install":

		if len(os.Args) < 3 {
			fmt.Println("Использование: wqa install <file>")
			return
		}

		err := installer.Install(os.Args[2])

		if err != nil {
			fmt.Println("Ошибка:", err)
		}

	case "update":

		if len(os.Args) < 3 {
			fmt.Println("Использование: wqa update <app>")
			return
		}

		err := updater.Update(os.Args[2])

		if err != nil {
			fmt.Println("Ошибка:", err)
		}

	case "list":

		err := list.Show()

		if err != nil {
			fmt.Println("Ошибка:", err)
		}

	case "remove":

		if len(os.Args) < 3 {
			fmt.Println(
				"Использование: wqa remove <id>",
			)
			return
		}

		err := remove.Remove(
			os.Args[2],
		)

		if err != nil {
			fmt.Println(
				"[ERROR]",
				err,
			)
		}

	case "run":

		if len(os.Args) < 3 {
			fmt.Println("Использование: wqa run <app|file>")
			return
		}

		target := os.Args[2]

		err := run.Run(target)

		if err != nil {
			fmt.Println("[ERROR]", err)
		}

	case "setup":
		err := setup.Setup()

		if err != nil {
			fmt.Println("[ERROR]", err)
		}

	default:
		fmt.Println("Неизвестная команда:", os.Args[1])
	}
}

func createProject(name string) {

	err := os.MkdirAll(
		filepath.Join(name, "app"),
		0755,
	)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(
		filepath.Join(name, "resources"),
		0755,
	)
	if err != nil {
		panic(err)
	}

	manifest := Manifest{
		WQA:       "1.0",
		ID:        "windroid." + name,
		Name:      name,
		Version:   "1.0.0",
		Publisher: "Win Studio",
		Runtime:   "wqbc",
		Entry:     "app/main.wqa",
		Architecture: []string{
			"x86_64",
			"arm64",
		},
	}

	data, _ := json.MarshalIndent(
		manifest,
		"",
		"  ",
	)

	os.WriteFile(
		filepath.Join(name, "wqa.json"),
		data,
		0644,
	)

	os.WriteFile(
		filepath.Join(name, "app", "main.wqa"),
		[]byte("// WQA Application\n"),
		0644,
	)

	fmt.Println("[OK] WQA проект создан:", name)
}

func help() {
	fmt.Println(`
WQA CLI 1.0.0
WinDroid Quick Application Package

Использование:

  wqa <команда> [аргументы]

Основные команды:

  init <name>
      Создать новый WQA проект

  build
      Собрать проект в .wqa пакет

  info <file>
      Показать информацию о WQA пакете

  verify <file>
      Проверить целостность пакета

  install <file>
      Установить WQA приложение

  list
      Показать установленные приложения

Управление:

  remove <id>
      Удалить приложение

  run <id>
      Запустить приложение

Дополнительно:

  --version
      Показать версию WQA CLI

  --help
      Показать эту справку

Пример:

  wqa init Calculator
  wqa build
  wqa install Calculator.wqa`)
}
