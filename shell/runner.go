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
			if s.checkParts(parts, 2) == false {
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
			if s.checkParts(parts, 3) == false {
				break
			}
			err := s.Store.Set(parts[1], parts[2])
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Print(">>> DONE \n")
			break
		case "delete", "del", "remove", "rm":
			// Handle delete command, ensure there is a key
			if s.checkParts(parts, 2) == false {
				break
			}
			err := s.Store.Delete(parts[1])
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Print(">>> DONE \n")
			break
		case "ls", "list", "keys":
			// Handle list command
			results, err := s.Store.List()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Print(">>> \n")
			fmt.Printf("%20v %20v\n", "Key", "Value")
			for _, result := range results {
				fmt.Printf("%20v %20v\n", result.Key, result.Value)
			}
			break
		default:
			fmt.Println("Invalid command")
			s.PrintUsuage()
		}

	}
}

func (s *Shell) checkParts(parts []string, size int) bool {
	if len(parts) != size {
		fmt.Println("Invalid command")
		s.PrintUsuage()
		return false
	}
	return true
}
