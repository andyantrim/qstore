package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andyantrim/qstore/store"
)

type Shell struct {
	History [10]string
	reader  *bufio.Reader
	Store   store.Store
}

func NewShell(store store.Store) *Shell {
	return &Shell{
		History: [10]string{},
		reader:  bufio.NewReader(os.Stdin),
		Store:   store,
	}
}

func (s *Shell) getInput() ([]string, error) {
	text, err := s.reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	// Bum the prompt back down
	text = strings.ToLower(text)
	parts := strings.Split(text, " ")
	if len(parts) < 1 {
		return nil, fmt.Errorf("Invalid command")
	}

	for i := len(s.History) - 1; i > 0; i-- {
		s.History[i] = s.History[i-1]
	}
	// Add to history if it passes checks
	s.History[0] = text

	return parts, err
}

func (s *Shell) PrintUsuage() {
	fmt.Println("Usuage: <command> <args>")
	fmt.Println("Example: get foo")
	fmt.Println("Example: set foo bar")
}

func (s *Shell) PrintHeader() {
	fmt.Println("Welcome to the shell")
	fmt.Println("Type '?' or 'help' for help")
	fmt.Println("Type 'exit', 'quit', or 'q' to quit")
}

func (s *Shell) PrintHistory() {
	for i := len(s.History) - 1; i >= 0; i-- {
		fmt.Printf("%d: %v\n", i, s.History[i])
	}
}
