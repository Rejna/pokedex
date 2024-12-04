package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	cmds := commands()
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for k, v := range cmds {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	fmt.Println("")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {
	cmds := commands()
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "
	for fmt.Print(prompt); scanner.Scan(); fmt.Print(prompt) {
		line := scanner.Text()
		cmd, ok := cmds[line]
		if ok {
			if err := cmd.callback(); err != nil {
				fmt.Errorf("error executing %s: %w", line, err)
			}
		} else {
			fmt.Printf("unknown command %s\n", line)
		}

	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
