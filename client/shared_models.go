package intigriti

type ProgramType struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type ConfidentialityLevel struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Status struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type WebLinks struct {
	Detail string `json:"detail"`
}

type TestingRequirements struct {
	IntigritiMe      bool   `json:"intigritiMe"`
	AutomatedTooling int    `json:"automatedTooling"`
	UserAgent        string `json:"userAgent"`
	RequestHeader    any    `json:"requestHeader"`
}
