package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func getMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"mapf": {
			name:        "mapf",
			description: "Displays the names of 20 locations in the Pokemon world.",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 locations in the Pokemon world.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Accepts an argument of a location in the pokemon world and displays pokemon to be found in the area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Get information about a pokemon you have captured.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all pokemon you have caught.",
			callback:    commandPokedex,
		},
	}
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	locationsPokemon, err := cfg.pokeapiClient.GetAllPokemonByArea(name)
	if err != nil {
		return err
	}

	for _, encounter := range locationsPokemon.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandMapf(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.GetLocations(cfg.next)
	if err != nil {
		return err
	}

	for _, locale := range locations.Results {
		fmt.Println(locale.Name)
	}
	cfg.next = locations.Next
	cfg.previous = locations.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}
	locations, err := cfg.pokeapiClient.GetLocations(cfg.previous)
	if err != nil {
		return err
	}

	for _, locale := range locations.Results {
		fmt.Println(locale.Name)
	}
	cfg.next = locations.Next
	cfg.previous = locations.Previous
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	if catchChance := rand.Intn(400); catchChance < pokemon.BaseExperience {
		fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")
	fmt.Printf("Successfully caught %s!\n", pokemon.Name)
	cfg.pokedex.UpdatePokedex(pokemon)
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.pokedex.InspectPokemon(pokemonName)
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, val := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types: ")
	for _, val := range pokemon.Types {
		fmt.Printf(" -%s\n", val.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex.GetPokedex() {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	commands := getMap()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Available commands include: ")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
