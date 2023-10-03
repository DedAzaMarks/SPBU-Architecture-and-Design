package commands

import (
	"bufio"
	"fmt"
	"github.com/DedAzaMarks/SPBU-Architecture-and-Design/projects/bash/internal/state"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Cat(s *state.State, filename string) (string, error) {
	if filename == "" {
		if s.PrevCommandOutput == "" {
			return "", fmt.Errorf("Usage: cat [FILE]")
		}
		content := s.PrevCommandOutput
		s.PrevCommandOutput = ""
		return content, nil
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("File not found %s", filename)
	}

	s.PrevCommandOutput = string(content)
	return string(content), nil
}

func Wc(s *state.State, filename string) (string, error) {
	var input io.Reader

	if filename == "" {
		if s.PrevCommandOutput == "" {
			return "", fmt.Errorf("Usage: wc [FILE]")
		}
		input = strings.NewReader(s.PrevCommandOutput)
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return "", fmt.Errorf("File not found %s", filename)
		}
		defer file.Close()
		input = file
	}
	content, err := wc(input)
	s.PrevCommandOutput = content
	return content, err
}

func wc(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)
	lineCount, wordCount, byteCount := 0, 0, 0

	for scanner.Scan() {
		lineCount++
		wordCount += len(strings.Fields(scanner.Text()))
		byteCount += len(scanner.Text()) + 1 // +1 for the newline character
	}

	return fmt.Sprintf("Lines: %d, Words: %d, Bytes: %d", lineCount, wordCount, byteCount), nil
}

func Pwd(s *state.State) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Error getting current directory: %v", err)
	}
	s.PrevCommandOutput = dir
	return dir, nil
}

func Echo(s *state.State, arg string) string {
	s.PrevCommandOutput = arg
	return arg
}
