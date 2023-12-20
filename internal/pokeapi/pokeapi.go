package pokeapi

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const endpoint = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{httpClient: http.Client{Timeout: time.Minute}}
}


func buildUrl(path string, offset, limit int) string {
	u, _ := url.Parse(endpoint + path)
	query := url.Values{}
	query.Set("limit", strconv.Itoa(limit))
	if offset != 0 {
		query.Set("offset", strconv.Itoa(offset))
	}

	u.RawQuery = query.Encode()

	return u.String()
}
