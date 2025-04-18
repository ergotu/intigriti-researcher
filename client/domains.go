package intigriti

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Domains struct {
	ID        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	Content   []struct {
		ID          string      `json:"id"`
		Type        ProgramType `json:"type"`
		Endpoint    string      `json:"endpoint"`
		Tier        Tier        `json:"tier"`
		Description string      `json:"description"`
	} `json:"content"`
}

type Tier struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func (c *Client) GetProgramDomains(programID string, version string) (Domains, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/programs/%s/domains/%s", c.BaseURL, programID, version), nil)
	if err != nil {
		return Domains{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return Domains{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Domains{}, fmt.Errorf("failed to fetch program domains: %s", resp.Status)
	}

	var domains Domains
	if err := json.NewDecoder(resp.Body).Decode(&domains); err != nil {
		return Domains{}, err
	}

	return domains, nil
}
