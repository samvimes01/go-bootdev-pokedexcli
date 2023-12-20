package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []area  `json:"results"`
}

func (c *Client) GetLocationAreas(offset, limit int) (LocationAreaResp, error) {
	url := buildUrl("location-area", offset, limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaResp{}, fmt.Errorf("Response failed with status code: %d and\n", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	areas := LocationAreaResp{}
	err = json.Unmarshal(body, &areas)

	if err != nil {
		return LocationAreaResp{}, fmt.Errorf("bad response JSON")
	}
	return areas, nil
}
