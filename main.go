package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Rejna/pokedex/internal/pokeapi"
	"github.com/Rejna/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
	Ch       *pokecache.Cache
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
		"map": {
			name:        "map",
			description: "List next 20 available locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 available locations",
			callback:    commandMapB,
		},
	}
}

func commandMap(c *config) error {
	if c.Next == "" {
		return errors.New("no more locations")
	}

	locationArea, err := pokeapi.GetLocationArea(c.Next, *c.Ch)
	if err != nil {
		return err
	}
	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	if locationArea.Previous != nil {
		c.Previous = *locationArea.Previous
	} else {
		c.Previous = ""
	}
	if locationArea.Next != nil {
		c.Next = *locationArea.Next
	} else {
		c.Next = ""
	}

	return nil
}

func commandMapB(c *config) error {
	if c.Previous == "" {
		return errors.New("no more locations")
	}

	locationArea, err := pokeapi.GetLocationArea(c.Previous, *c.Ch)
	if err != nil {
		return err
	}
	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	if locationArea.Previous != nil {
		c.Previous = *locationArea.Previous
	} else {
		c.Previous = ""
	}
	if locationArea.Next != nil {
		c.Next = *locationArea.Next
	} else {
		c.Next = ""
	}

	return nil
}

func commandHelp(c *config) error {
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

func commandExit(c *config) error {
	os.Exit(0)
	return nil
}

func main() {
	cmds := commands()
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "
	cache := pokecache.NewCache(time.Second * 10)
	config := config{Next: "https://pokeapi.co/api/v2/location-area", Ch: &cache}

	for fmt.Print(prompt); scanner.Scan(); fmt.Print(prompt) {
		line := scanner.Text()
		cmd, ok := cmds[line]
		if ok {
			if err := cmd.callback(&config); err != nil {
				fmt.Println(fmt.Errorf("error executing %s: %w", line, err))
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
