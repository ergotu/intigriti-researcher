package intigriti

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProgramDetail struct {
	ID                   string               `json:"id"`
	Handle               string               `json:"handle"`
	Name                 string               `json:"name"`
	Following            bool                 `json:"following"`
	ConfidentialityLevel ConfidentialityLevel `json:"confidentialityLevel"`
	Status               Status               `json:"status"`
	Type                 ProgramType          `json:"type"`
	Domains              Domains              `json:"domains"`
	RulesOfEngagement    RulesOfEngagement    `json:"rulesOfEngagement"`
	WebLinks             WebLinks             `json:"webLinks"`
}

func (c *Client) GetProgramDetail(programID string) (ProgramDetail, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/programs/%s", c.BaseURL, programID), nil)
	if err != nil {
		return ProgramDetail{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return ProgramDetail{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ProgramDetail{}, fmt.Errorf("failed to fetch program detail: %s", resp.Status)
	}

	var detail ProgramDetail
	if err := json.NewDecoder(resp.Body).Decode(&detail); err != nil {
		return ProgramDetail{}, err
	}

	return detail, nil
}
