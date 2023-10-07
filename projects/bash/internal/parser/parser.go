package parser

import (
	"strings"
)

type Command struct {
	Command   string   // Command - character sequence, serves as identifier of which function to call
	Arguments []string // Arguments - program arguments, doesn't have Command as first element
}

// ParseCommandLine - accepts string and splits it into sequence of commands separated with '|'.
func ParseCommandLine(s string) ([]Command, error) {
	var res []Command
	chunks := strings.Split(s, "|") // todo - cmd || cmd будет страдать, се ля ви, потом переделаем
	for _, chunk := range chunks {
		cmdArgs := strings.Split(strings.TrimSpace(chunk), " ")
		res = append(res, Command{Command: cmdArgs[0], Arguments: cmdArgs[1:]})
	}
	return res, nil
}
