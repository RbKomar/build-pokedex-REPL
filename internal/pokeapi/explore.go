package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) unmarshalExploreResponse(response []byte) (ResponseShallowExplore, error) {
	exploreResponse := ResponseShallowExplore{}
	if err := json.Unmarshal(response, &exploreResponse); err != nil {
		return ResponseShallowExplore{}, err
	}
	return exploreResponse, nil
}

func (c *Client) ExploreLocation(locationName string) (ResponseShallowExplore, error) {
	url := baseURL + "/location-area/" + locationName

	if response, exists := c.cache.Get(url); exists {
		log.Println("Hit cache, instant response")
		return c.unmarshalExploreResponse(response)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseShallowExplore{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseShallowExplore{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseShallowExplore{}, err
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	c.cache.Add(url, body)
	return c.unmarshalExploreResponse(body)
}
