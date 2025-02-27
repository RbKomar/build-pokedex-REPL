package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) unmarshalLocationResponse(response []byte) (ResponseShallowLocations, error) {
	locationResponse := ResponseShallowLocations{}
	if err := json.Unmarshal(response, &locationResponse); err != nil {
		return ResponseShallowLocations{}, err
	}
	return locationResponse, nil
}

func (c *Client) ListLocations(pageUrl *string) (ResponseShallowLocations, error) {
	url := locationUrl
	if pageUrl != nil {
		url = *pageUrl
	}

	if response, exists := c.cache.Get(url); exists {
		log.Println("Hit cache, instant response")
		return c.unmarshalLocationResponse(response)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	c.cache.Add(url, body)
	return c.unmarshalLocationResponse(body)
}
