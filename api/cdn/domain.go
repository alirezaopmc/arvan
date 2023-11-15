package cdn

import (
	"encoding/json"
	"fmt"
)

func (d *API) GetDomains(search string, perPage, page int) ([]Domain, error) {
	url := fmt.Sprintf("%s/domains", d.BaseURL)

	params := map[string]string{
		"search":   search,
		"per_page": fmt.Sprintf("%d", perPage),
		"page":     fmt.Sprintf("%d", page),
	}

	responseBody, err := d.HttpProxy.GET(url, params)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data  []Domain `json:"data"`
		Links struct {
			First string `json:"first"`
			Last  string `json:"last"`
			Prev  string `json:"prev"`
			Next  string `json:"next"`
		} `json:"links"`
		Meta struct {
			CurrentPage int    `json:"current_page"`
			From        int    `json:"from"`
			LastPage    int    `json:"last_page"`
			Path        string `json:"path"`
			PerPage     int    `json:"per_page"`
			To          int    `json:"to"`
			Total       int    `json:"total"`
		} `json:"meta"`
	}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
