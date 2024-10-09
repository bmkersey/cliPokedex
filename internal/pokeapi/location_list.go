package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationsResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResp{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResp{}, err
	}

	locations := LocationsResp{}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsResp{}, nil
	}

	return locations, nil

}
