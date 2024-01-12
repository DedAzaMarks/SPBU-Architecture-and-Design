package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/DedAzaMarks/SPBU-Architecture-and-Design/projects/bash/internal/parser"
	"github.com/DedAzaMarks/SPBU-Architecture-and-Design/projects/bash/internal/state"
)

func main() {
	// todo - делать что то с глобальным стейтом. он как минимум понадобится на моменте исполнения
	s := state.NewState()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		cmd := scanner.Text()
		commands, err := parser.ParseCommandLine(cmd)
		if err != nil {
			log.Printf("parsing error: %v", err)
		}
		if err := s.EvaluateCommands(commands); err != nil {
			log.Printf("error: %v", err)
		}
		fmt.Print("> ")
	}
}
