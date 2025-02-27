package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) unmarshalPokemonResponse(response []byte) (Pokemon, error) {
	pokemonResponse := Pokemon{}
	if err := json.Unmarshal(response, &pokemonResponse); err != nil {
		return Pokemon{}, err
	}
	return pokemonResponse, nil
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := pokemonUrl + name

	if response, exists := c.cache.Get(url); exists {
		log.Println("Hit cache, instant response")
		return c.unmarshalPokemonResponse(response)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	c.cache.Add(url, body)
	return c.unmarshalPokemonResponse(body)

}
