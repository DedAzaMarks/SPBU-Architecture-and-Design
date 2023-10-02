package state

import "os"

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

func (s *State) GetState() State {
	return *s
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
