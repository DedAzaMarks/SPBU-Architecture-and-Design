package state

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Cat(s *State, filenames []string) (string, error) {
	if len(filenames) == 0 {
		if s.PrevCommandOutput == "" {
			return "", fmt.Errorf("Usage: cat [FILE]")
		}
		content := s.PrevCommandOutput
		s.PrevCommandOutput = ""
		return content, nil
	}

	var content string
	for _, filename := range filenames {
		buf, err := os.ReadFile(filename)
		if err != nil {
			return "", fmt.Errorf("File not found %s", filename)
		}
		content += string(buf)
	}

	s.PrevCommandOutput = content
	return content, nil
}

func Wc(s *State, filename []string) (string, error) {
	var input io.Reader

	if len(filename) == 0 {
		if s.PrevCommandOutput == "" {
			return "", fmt.Errorf("Usage: wc [FILE]")
		}
		input = strings.NewReader(s.PrevCommandOutput)
	} else {
		file, err := os.Open(filename[0])
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

func Pwd(s *State, args []string) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf("pwd: too many arguments")
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Error getting current directory: %v", err)
	}
	s.PrevCommandOutput = dir
	return dir, nil
}

func Echo(s *State, args []string) (string, error) {
	var builder strings.Builder
	for _, arg := range args {
		builder.WriteString(arg)
	}
	s.PrevCommandOutput = builder.String()
	return s.PrevCommandOutput, nil
}
