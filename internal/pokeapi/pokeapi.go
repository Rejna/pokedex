package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Rejna/pokedex/internal/pokecache"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(url string, c pokecache.Cache) (LocationAreaResponse, error) {
	data, err := sendGetApiRequest(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error while getting location area data from API: %w", err)
	}
	var locationArea LocationAreaResponse
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error parsing JSON body: %w", err)
	}

	return locationArea, nil
}

func sendGetApiRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "pokedex-go")
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}
	var httpClient http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bad response: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}
	return bodyBytes, nil
}
