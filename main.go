package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/stephenoveson/pokedexcli/pokeapi"
	"github.com/stephenoveson/pokedexcli/pokedex"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       pokedex.Pokedex
	next          *string
	previous      *string
}

func cleanInput(s string) []string {
	output := strings.ToLower(s)
	return strings.Fields(output)
}

func main() {
	commands := getMap()
	config := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, time.Minute*5),
		pokedex:       pokedex.InitPokedex(),
	}
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := cleanInput(reader.Text())

		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		command, exists := commands[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
