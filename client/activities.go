package intigriti

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	ActivityDetail struct{}
	ActivityType   struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	}
)

type Activity struct {
	ProgramID string         `json:"programId"`
	Activity  ActivityDetail `json:"activity"`
	Type      ActivityType   `json:"type"`
	CreatedAt int            `json:"createdAt"`
	Following bool           `json:"following"`
}

func (c *Client) GetActivities() ([]Activity, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/programs/activities", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch program detail: %s", resp.Status)
	}

	var result struct {
		MaxCount int        `json:"maxCount"`
		Records  []Activity `json:"records"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Records, nil
}
