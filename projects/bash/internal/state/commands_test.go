package state

import (
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	tests := []struct {
		name           string
		state          *State
		filename       []string
		expectedOutput string
		expectedError  string
	}{
		{
			name:           "Read from file",
			state:          &State{},
			filename:       []string{"test.txt"}, // Provide the actual file path if needed
			expectedOutput: "Hello, World!",
			expectedError:  "",
		},
		{
			name:           "Read from previous command output",
			state:          &State{PrevCommandOutput: "Previous Output"},
			filename:       []string{""},
			expectedOutput: "Previous Output",
			expectedError:  "",
		},
		{
			name:           "No input provided",
			state:          &State{},
			filename:       []string{""},
			expectedOutput: "",
			expectedError:  "Usage: cat [FILE]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := Cat(tt.state, tt.filename)

			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("Cat() error = %v", err.Error())
				return
			}

			if output != tt.expectedOutput {
				t.Errorf("Cat() error = %v", err.Error())
			}
		})
	}
}

func TestWc(t *testing.T) {
	tests := []struct {
		name           string
		state          *State
		filename       []string
		expectedOutput string
		expectedError  string
	}{
		{
			name:           "Read from file",
			state:          &State{},
			filename:       []string{"commands.go"}, // Provide the actual file path if needed
			expectedOutput: "Lines: 2, Words: 4, Bytes: 24",
			expectedError:  "",
		},
		{
			name:           "Read from previous command output",
			state:          &State{PrevCommandOutput: "Previous Output"},
			filename:       []string{""},
			expectedOutput: "Lines: 1, Words: 2, Bytes: 16",
			expectedError:  "",
		},
		{
			name:           "No input provided",
			state:          &State{},
			filename:       []string{""},
			expectedOutput: "",
			expectedError:  "Usage: wc [FILE]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := Wc(tt.state, tt.filename)

			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("Wc() error = %v, expected error = %v", err, tt.expectedError)
				return
			}

			if output != tt.expectedOutput {
				t.Errorf("Wc() output = %v, expected output = %v", output, tt.expectedOutput)
			}
		})
	}
}

func Test_wc(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectedError  string
	}{
		{
			name:           "Test wc with input",
			input:          "Line 1\nLine 2\nLine 3",
			expectedOutput: "Lines: 3, Words: 6, Bytes: 21",
			expectedError:  "",
		},
		{
			name:           "Test wc with empty input",
			input:          "",
			expectedOutput: "Lines: 0, Words: 0, Bytes: 0",
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := wc(strings.NewReader(tt.input))

			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("wc() error = %v, expected error = %v", err, tt.expectedError)
				return
			}

			if output != tt.expectedOutput {
				t.Errorf("wc() output = %v, expected output = %v", output, tt.expectedOutput)
			}
		})
	}
}
