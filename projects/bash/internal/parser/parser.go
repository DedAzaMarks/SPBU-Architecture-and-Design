package parser

import (
	"strings"
)

type Command struct {
	Command   string
	Arguments []string
}

func ParseCommandLine(s string) ([]Command, error) {
	var res []Command
	chunks := strings.Split(s, "|") // todo - cmd || cmd будет страдать, се ля ви, потом переделаем
	for _, chunk := range chunks {
		cmdArgs := strings.Split(strings.TrimSpace(chunk), " ")
		res = append(res, Command{Command: cmdArgs[0], Arguments: cmdArgs[1:]})
	}
	return res, nil
}
