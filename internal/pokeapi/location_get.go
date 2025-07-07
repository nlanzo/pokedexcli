package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetLocation(location string) (ResponseLocation, error) {
	url := baseURL + "/location-area/" + location


	if cacheData, ok := c.cache.Get(url); ok {
		locationRes := ResponseLocation{}
		err := json.Unmarshal(cacheData, &locationRes)
		if err != nil {
			return ResponseLocation{}, err
		}
		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocation{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocation{}, err
	}

	locationRes := ResponseLocation{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return ResponseLocation{}, err
	}

	c.cache.Add(url, data)

	return locationRes, nil
}