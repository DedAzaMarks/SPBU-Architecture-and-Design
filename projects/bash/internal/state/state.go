package state

import (
	"fmt"
	"os"

	"github.com/DedAzaMarks/SPBU-Architecture-and-Design/projects/bash/internal/parser"
)

type State struct {
	availableCommands map[string]struct{}

	GlobalVariables map[string]string

	CommandContent string

	PrevCommandOutput string
	PrevReturnCode    int
}

func (s *State) CheckCommand(cmd string) bool {
	_, ok := s.availableCommands[cmd]
	return ok
}

func (s *State) EvaluateCommands(commands []parser.Command) error {
	for _, command := range commands {
		cmd, err := s.substituteVariables(command)
		if err != nil {
			return fmt.Errorf("variable substitution error: %w", err)
		}
		_ = cmd
	}
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

func NewState() *State {
	return &State{
		availableCommands: map[string]struct{}{
			"cat":  {},
			"echo": {},
			"wc":   {},
			"pwd":  {},
			"exit": {},
		},
		GlobalVariables: map[string]string{
			"HOME": os.Getenv("HOME"),
			"USER": os.Getenv("USER"),
		},
	}
}
