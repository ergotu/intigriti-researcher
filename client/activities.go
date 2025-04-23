package intigriti

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	CreatedAt int64          `json:"createdAt"`
	Following bool           `json:"following"`
}

// GetActivitiesOptions defines the optional parameters for the GetActivities method
type GetActivitiesOptions struct {
	CreatedSince *int64
	Following    *bool
	Limit        *int
	Offset       *int
}

func (c *Client) GetActivities(opts GetActivitiesOptions) ([]Activity, error) {
	baseURL := fmt.Sprintf("%s/v1/programs/activities", c.BaseURL)
	params := url.Values{}

	if opts.CreatedSince != nil {
		params.Add("createdSince", fmt.Sprintf("%d", *opts.CreatedSince))
	}
	if opts.Following != nil {
		params.Add("following", fmt.Sprintf("%t", *opts.Following))
	}
	if opts.Limit != nil {
		params.Add("limit", fmt.Sprintf("%d", *opts.Limit))
	}
	if opts.Offset != nil {
		params.Add("offset", fmt.Sprintf("%d", *opts.Offset))
	}

	fullURL := baseURL
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", fullURL, nil)
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
		return nil, fmt.Errorf("failed to fetch activities: %s", resp.Status)
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
