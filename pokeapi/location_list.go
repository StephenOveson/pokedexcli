package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageUrl *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResponse := LocationAreas{}
		err := json.Unmarshal(val, &locationsResponse)
		if err != nil {
			return LocationAreas{}, err
		}

		return locationsResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, nil
	}

	locationsResponse := LocationAreas{}
	err = json.Unmarshal(dat, &locationsResponse)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(url, dat)
	return locationsResponse, nil
}
