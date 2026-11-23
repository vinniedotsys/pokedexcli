package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/vinniedotsys/pokedexcli/internal/pokecache"
	"time"
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
	var cache pokecache.Cache
	cache.NewCache(5 * time.Minute)
	
	data, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(data, &result)
		return result, err

	}

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
	cache.Add(url, body)

	return result, err

}
