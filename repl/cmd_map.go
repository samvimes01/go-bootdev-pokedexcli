package repl

import (
	"fmt"

	"github.com/samvimes01/go-bootdev-pokedexcli/internal/pokeapi"
	"github.com/samvimes01/go-bootdev-pokedexcli/utils"
)

type area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getAreas(cfg *replConfig, offset int) error {
	var areas pokeapi.LocationAreaResp
	var err error
	if cache, ok := cfg.Cache.Get(offset); ok {
		areas = cache
	} else {
		areas, err = cfg.PokeapiClient.GetLocationAreas(offset, cfg.Limit)
		if err != nil {
			return err
		}
	}

	cfg.PreviousOffset = utils.GetOffsetFromUrl(areas.Previous)
	cfg.NextOffset = utils.GetOffsetFromUrl(areas.Next)

	fmt.Println("Location areas:")
	for _, area := range areas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}

func commandMap(cfg *replConfig) error {
	err := getAreas(cfg, cfg.NextOffset)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(cfg *replConfig) error {
	if cfg.PreviousOffset == 0 && cfg.NextOffset == cfg.Limit {
		return fmt.Errorf("no previous page")
	}
	err := getAreas(cfg, cfg.PreviousOffset)
	if err != nil {
		return err
	}
	return nil
}
