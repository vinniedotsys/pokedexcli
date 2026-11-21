package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



type LocationAreas struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreas, error) {
	var result LocationAreas
	res, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return result, 	fmt.Errorf("HTTP Error: %v", res.StatusCode)
	}

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err

}
