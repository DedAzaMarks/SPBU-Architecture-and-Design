package state

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/DedAzaMarks/SPBU-Architecture-and-Design/projects/bash/internal/parser"
)

// State - holds information on current session variables and commands output
type State struct {
	availableCommands map[string]func(state *State, args []string) (string, error)

	// GlobalVariables - holds key-value pairs. Key is a variable name, Value is a variable value
	GlobalVariables map[string]string

	// CommandContent - holds full command from stdin
	CommandContent string

	// PrevCommandOutput - holds output of the previous command in pipeline sequence.
	// Should be cleaned with State.Reset method after all commands in pipeline are finished.
	PrevCommandOutput string

	// PrevReturnCode - exit code of the previous command
	PrevReturnCode int
}

// CheckCommand - is cmd present in default commands set
func (s *State) CheckCommand(cmd string) bool {
	_, ok := s.availableCommands[cmd]
	return ok
}

// EvaluateCommands - evaluate sequence of commands
func (s *State) EvaluateCommands(commands []parser.Command) error {
	for _, command := range commands {
		cmd, err := s.substituteVariables(command)
		if err != nil {
			return fmt.Errorf("variable substitution error: %w", err)
		}
		if s.CheckCommand(cmd.Command) {
			_, err := s.availableCommands[cmd.Command](s, cmd.Arguments)
			if err != nil {
				return fmt.Errorf("evaluation error: %w", err)
			}
			continue
		}
		// todo - обработка переменных
		output, err := exec.Command(cmd.Command, cmd.Arguments...).CombinedOutput()
		if err != nil {
			return fmt.Errorf("exec error: %w", err)
		}
		s.PrevCommandOutput = string(output)
	}
	fmt.Println(s.PrevCommandOutput)
	s.Reset()
	return nil
}

func (s *State) substituteVariables(command parser.Command) (parser.Command, error) {
	newCommand := s.substituteVariable(command.Command)
	newArguments := make([]string, 0, len(command.Arguments))
	for _, arg := range command.Arguments {
		newArguments = append(newArguments, s.substituteVariable(arg))
	}
	return parser.Command{Command: newCommand, Arguments: newArguments}, nil
}

func (s *State) substituteVariable(word string) string {
	var newWord []byte
	for i := 0; i < len(word); i++ {
		if word[i] == '$' {
			var varName []byte
			for ; i < len(word) && word[i] != '$'; i++ {
				varName = append(varName, word[i])
			}
			gVar := s.GlobalVariables[string(varName)]
			newWord = append(newWord, gVar...)
		} else {
			newWord = append(newWord, word[i])
		}
	}
	return string(newWord)
}

// Reset - reset State's fields for next command input
func (s *State) Reset() {
	s.PrevCommandOutput = s.PrevCommandOutput[:0]
	s.CommandContent = s.CommandContent[:0]
}

// NewState - creates new state, which holds default commands and some basic environment variables
func NewState() *State {
	return &State{
		availableCommands: map[string]func(state *State, strings []string) (string, error){
			"cat":  Cat,
			"echo": Echo,
			"wc":   Wc,
			"pwd":  Pwd,
			"grep": Grep,
			"exit": func(state *State, strings []string) (string, error) {
				os.Exit(state.PrevReturnCode)
				return "", nil
			},
		},
		GlobalVariables: map[string]string{
			"HOME": os.Getenv("HOME"),
			"USER": os.Getenv("USER"),
		},
	}
}
