package pokeapi

type LocationAreas struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaPokemon struct {
	Location          Results             `json:"location"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Results `json:"pokemon"`
}
