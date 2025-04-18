package intigriti

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RulesOfEngagement struct {
	Attachments []any  `json:"attachments"`
	ID          string `json:"id"`
	CreatedAt   int    `json:"createdAt"`
	Content     struct {
		Description         string              `json:"description"`
		TestingRequirements TestingRequirements `json:"testingRequirements"`
		SafeHarbour         bool                `json:"safeHarbour"`
	} `json:"content"`
}

func (c *Client) GetROE(programID string, version string) (RulesOfEngagement, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/programs/%s/rules-of-engagements/%s", c.BaseURL, programID, version), nil)
	if err != nil {
		return RulesOfEngagement{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return RulesOfEngagement{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RulesOfEngagement{}, fmt.Errorf("failed to fetch program rules of engagement: %s", resp.Status)
	}

	var rulesOfEngagement RulesOfEngagement
	if err := json.NewDecoder(resp.Body).Decode(&rulesOfEngagement); err != nil {
		return RulesOfEngagement{}, err
	}

	return rulesOfEngagement, nil
}
