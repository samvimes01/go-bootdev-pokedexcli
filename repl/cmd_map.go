package repl

import (
	"fmt"
	"log"

	"github.com/samvimes01/go-bootdev-pokedexcli/utils"
)

type area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getAreas(cfg *replConfig, offset int) {
	areas, err := cfg.PokeapiClient.GetLocationAreas(offset, cfg.Limit)
	if err != nil {
		log.Fatal(err)
	}

	cfg.PreviousOffset = utils.GetOffsetFromUrl(areas.Previous)
	cfg.NextOffset = utils.GetOffsetFromUrl(areas.Next)

	fmt.Println("Location areas:")
	for _, area := range areas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
}

func commandMap(cfg *replConfig) error {
	getAreas(cfg, cfg.NextOffset)
	return nil
}

func commandMapb(cfg *replConfig) error {
	if cfg.PreviousOffset == 0 && cfg.NextOffset == cfg.Limit {
		return fmt.Errorf("no previous page")
	}
	getAreas(cfg, cfg.PreviousOffset)
	return nil
}
