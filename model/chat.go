package model

type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type AnalyzePayload struct {
	Inputs `json:"inputs"`
}

type ChatPayload struct {
	Inputs    string `json:"inputs"`
	MaxTokens int    `json:"max_tokens"`
}
