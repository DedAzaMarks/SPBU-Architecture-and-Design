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
	_ = s
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		cmd := scanner.Text()
		command, err := parser.ParseCommandLine(cmd)
		if err != nil {
			log.Printf("error : %v", err)
		}
		_ = command
		fmt.Print("> ")
	}
}
