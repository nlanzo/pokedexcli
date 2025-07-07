package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (ResponseLocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil && *pageURL != "" {
		url = *pageURL
	}


	if cacheData, ok := c.cache.Get(url); ok {
		locationsRes := ResponseLocationAreas{}
		err := json.Unmarshal(cacheData, &locationsRes)
		if err != nil {
			return ResponseLocationAreas{}, err
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocationAreas{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	locationsRes := ResponseLocationAreas{}
	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	c.cache.Add(url, data)

	return locationsRes, nil
}
