package pokedex

import (
	"github.com/stephenoveson/pokedexcli/pokeapi"
)

type Pokedex struct {
	pokedex map[string]pokeapi.Pokemon
}

func InitPokedex() Pokedex {

	return Pokedex{
		pokedex: map[string]pokeapi.Pokemon{},
	}
}

func (p Pokedex) UpdatePokedex(pokemon pokeapi.Pokemon) {
	p.pokedex[pokemon.Name] = pokemon
}

func (p Pokedex) InspectPokemon(key string) (pokeapi.Pokemon, bool) {
	entry, ok := p.pokedex[key]
	return entry, ok
}

func (p Pokedex) GetPokedex() map[string]pokeapi.Pokemon {
	return p.pokedex
}
