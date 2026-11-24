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


type SpecificArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []*string `json:"condition_values"`
				PokemonDetails  *string   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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



func GetSpecificArea(url string) (SpecificArea, error) {
	var result SpecificArea
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
