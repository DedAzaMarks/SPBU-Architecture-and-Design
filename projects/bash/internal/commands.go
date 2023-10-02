package internal

import (
	"bufio"
	"fmt"
	"github.com/lelouchhh/SPBU-Architecture-and-Design/task-1/projects/bash/internal/state"
	"io/ioutil"
	"os"
	"strings"
)

func Cat(filename string) (string, error) {
	state.
		content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func Wc(filename string) (int, int, int, string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка открытия файла: %v\n", err)
		return 0, 0, 0, ""
	}

	scanner := bufio.NewScanner(file)
	lineCount, wordCount, byteCount := 0, 0, 0

	for scanner.Scan() {
		lineCount++
		wordCount += len(strings.Fields(scanner.Text()))
		byteCount += len(scanner.Text()) + 1 // +1 для учета символа перевода строки
	}

	return lineCount, wordCount, byteCount, filename
}

func Pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения текущей директории: %v\n", err)
		return
	}
	fmt.Println(dir)
}

func Echo(args ...string) string {
	return strings.Join(args[:], ",")
}
