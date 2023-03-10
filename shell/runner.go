package shell

import (
	"fmt"
	"io"
)

func (s *Shell) Run() {
	s.PrintHeader()

	for {
		// Bump the prompt back down
		fmt.Println()

		// Read query
		fmt.Print(">> ")
		parts, err := s.getInput()

		// If there was an error, print it and continue
		if err != nil {
			if err == io.EOF {
				return
			}
			s.PrintUsuage()
			continue
		}

		switch parts[0] {
		case "?", "help":
			// Handle help command
			s.PrintUsuage()
			break
		case "exit", "quit", "q":
			// Handle exit command
			return
		case "history":
			// Handle history command
			s.PrintHistory()
			break
		case "get":
			// Handle get command, ensure there is a key
			if len(parts) < 2 {
				fmt.Println("Invalid command")
				s.PrintUsuage()
				break
			}
			results, err := s.Store.Get(parts[1])
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Printf(">>> %v\n", results)
			break
		case "set":
			// Handle set command, ensure there is a key and value
			if len(parts) < 3 {
				fmt.Println("Invalid command")
				s.PrintUsuage()
				break
			}
			err := s.Store.Set(parts[1], parts[2])
			if err != nil {
				fmt.Println(err.Error())
			}
			break
		default:
			fmt.Println("Invalid command")
			s.PrintUsuage()
		}

	}
}
