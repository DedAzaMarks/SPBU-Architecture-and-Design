package state

import (
	"bufio"
	"flag"
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

func Grep(s *State, args []string) (string, error) {
	grepFlags := flag.NewFlagSet("grep", flag.ExitOnError)
	var caseInsensitive, wholeWord bool
	var linesAfter int
	grepFlags.BoolVar(&caseInsensitive, "i", false, "case sensitivity")
	grepFlags.BoolVar(&wholeWord, "w", false, "whole word")
	grepFlags.IntVar(&linesAfter, "A", 0, "lines after match")
	if err := grepFlags.Parse(args); err != nil {
		return "", err
	}

	if len(grepFlags.Args()) < 1 {
		return "", fmt.Errorf("Usage: grep [-i] [-w] [-A num] pattern [file]")
	}
	pattern := grepFlags.Arg(0)
	var inputReader *os.File
	if len(grepFlags.Args()) >= 2 {
		filename := grepFlags.Arg(1)
		var err error
		inputReader, err = os.Open(filename)
		if err != nil {
			return "", err
		}
		defer inputReader.Close()
	} else {
		tempFile, err := os.CreateTemp("", "tempfile.txt")
		if err != nil {
			fmt.Println("Error creating temporary file:", err)
			return "", err
		}
		defer tempFile.Close()

		// Write the string content to the temporary file
		_, err = tempFile.WriteString(s.PrevCommandOutput)
		if err != nil {
			fmt.Println("Error writing to temporary file:", err)
			return "", err
		}
		inputReader, err = os.Open(tempFile.Name())
		if err != nil {
			return "", err
		}
	}
	res, err := grep(s, inputReader, caseInsensitive, wholeWord, pattern, linesAfter)
	if err != nil {
		return "", err
	}
	return res, nil
}

func grep(s *State, inputReader *os.File, caseInsensitive, wholeWord bool, pattern string, linesAfter int) (string, error) {
	var result []string
	scanner := bufio.NewScanner(inputReader)
	var buffer []string
	matching := false // Flag to track if a match occurred
	pattern = strings.Replace(pattern, "'", "", -1)
	pattern = strings.Replace(pattern, "\"", "", -1)
	if caseInsensitive {
		pattern = strings.ToLower(pattern)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if caseInsensitive {
			line = strings.ToLower(line)
		}

		matched := false
		if wholeWord {
			// Check for word boundaries before and after the pattern
			matched = DoesConsistWholeWord(line, pattern)
		} else {
			matched = strings.Contains(line, pattern)
		}

		if matched {

			matching = true
			if len(buffer) > 0 {
				result = append(result, buffer...)
				buffer = nil
			}
			result = append(result, line)
		} else if matching && linesAfter > 0 {
			buffer = append(buffer, line)
			if len(buffer) > linesAfter {
				buffer = buffer[1:]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if matching {
		result = append(result, buffer...)
	}

	s.PrevCommandOutput = strings.Join(result, "\n")
	return s.PrevCommandOutput, nil
}

func DoesConsistWholeWord(source, check string) bool {

	words := make(map[string]struct{})
	for _, w := range strings.Fields(strings.ToLower(source)) {
		words[w] = struct{}{}
	}

	_, ok := words[check]
	return ok
}
func parseInt(s string) int {
	i := 0
	for _, ch := range s {
		i = i*10 + int(ch-'0')
	}
	return i
}
