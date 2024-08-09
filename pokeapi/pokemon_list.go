package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetAllPokemonByArea(name string) (LocationAreaPokemon, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		LocationPokemonResponse := LocationAreaPokemon{}
		err := json.Unmarshal(val, &LocationPokemonResponse)
		if err != nil {
			return LocationAreaPokemon{}, nil
		}

		return LocationPokemonResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaPokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaPokemon{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaPokemon{}, nil
	}

	LocationPokemonResponse := LocationAreaPokemon{}
	err = json.Unmarshal(dat, &LocationPokemonResponse)
	if err != nil {
		return LocationAreaPokemon{}, err
	}

	c.cache.Add(url, dat)
	return LocationPokemonResponse, nil
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		LocationPokemonResponse := Pokemon{}
		err := json.Unmarshal(val, &LocationPokemonResponse)
		if err != nil {
			return Pokemon{}, nil
		}

		return LocationPokemonResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	LocationPokemonResponse := Pokemon{}
	err = json.Unmarshal(dat, &LocationPokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return LocationPokemonResponse, nil
}
