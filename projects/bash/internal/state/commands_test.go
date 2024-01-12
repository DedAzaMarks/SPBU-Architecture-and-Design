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
			filename:       nil,
			expectedOutput: "Previous Output",
			expectedError:  "",
		},
		{
			name:           "No input provided",
			state:          &State{},
			filename:       nil,
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
			filename:       []string{"test.txt"}, // Provide the actual file path if needed
			expectedOutput: "Lines: 1, Words: 2, Bytes: 14",
			expectedError:  "",
		},
		{
			name:           "Read from previous command output",
			state:          &State{PrevCommandOutput: "Previous Output"},
			filename:       nil,
			expectedOutput: "Lines: 1, Words: 2, Bytes: 16",
			expectedError:  "",
		},
		{
			name:           "No input provided",
			state:          &State{},
			filename:       nil,
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
func TestGrep(t *testing.T) {
	// Create a test State instance
	s := &State{}

	// Create a test input string
	input := `This is a Test line 1.
This is a test line 2.
This is a test line 3.
Pattern on line 4.
This is a test line 5.`

	// Define the test cases
	testCases := []struct {
		name           string
		args           []string
		expectedOutput string
	}{
		{
			name:           "Match entire line (case-insensitive)",
			args:           []string{"-i", "pattern"},
			expectedOutput: "Pattern on line 4.",
		},
		{
			name:           "Match entire line (case-sensitive)",
			args:           []string{"Pattern"},
			expectedOutput: "Pattern on line 4.",
		},
		{
			name:           "Match whole word (case-insensitive)",
			args:           []string{"-i", "-w", "is"},
			expectedOutput: "This is a Test line 1.\nThis is a test line 2.\nThis is a test line 3.\nThis is a test line 5.",
		},
		{
			name:           "Match whole word (case-sensitive)",
			args:           []string{"-w", "is"},
			expectedOutput: "This is a Test line 1.\nThis is a test line 2.\nThis is a test line 3.\nThis is a test line 5.", // No match because it's case-sensitive
		},
		{
			name:           "Match with lines after (case-insensitive)",
			args:           []string{"-i", "-A", "2", "test"},
			expectedOutput: "This is a Test line 1.\nThis is a test line 2.\nThis is a test line 3.\nPattern on line 4.\nThis is a test line 5.",
		},
		{
			name:           "Match with lines after (case-sensitive)",
			args:           []string{"-A", "2", "Test"},
			expectedOutput: "This is a Test line 1.\nThis is a test line 2.\nThis is a test line 3.",
		},
	}

	// Run the test cases
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if i == 2 {
				print()
			}
			// Set the PrevCommandOutput of the State to the test input
			s.PrevCommandOutput = input

			output, err := Grep(s, tc.args)
			if err != nil {
				t.Errorf("Grep returned an error: %v", err)
			}
			if output != tc.expectedOutput {
				t.Errorf("Expected output: %s, but got: %s", tc.expectedOutput, output)
			}
		})
	}
}
