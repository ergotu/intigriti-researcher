package intigriti

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Program struct {
	ID                   string               `json:"id"`
	Handle               string               `json:"handle"`
	Name                 string               `json:"name"`
	Following            bool                 `json:"following"`
	MinBounty            MinBounty            `json:"minBounty"`
	MaxBounty            MaxBounty            `json:"maxBounty"`
	ConfidentialityLevel ConfidentialityLevel `json:"confidentialityLevel"`
	Status               Status               `json:"status"`
	Type                 ProgramType          `json:"type"`
	WebLinks             WebLinks             `json:"webLinks"`
}

type MinBounty struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

type MaxBounty struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

func (c *Client) GetPrograms() ([]Program, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/v1/programs", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch programs: %s, body: %s", resp.Status, body)
	}

	var result struct {
		MaxCount int       `json:"maxCount"`
		Records  []Program `json:"records"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err) // Log the erro
		return nil, err
	}

	return result.Records, nil
}
