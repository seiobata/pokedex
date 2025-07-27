package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(pageURL *string) (LocationArea, error) {
	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locations := LocationArea{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationArea{}, err
	}
	return locations, nil
}
