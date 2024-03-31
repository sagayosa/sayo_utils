package aitypes

// chat/completions
type CompletionsResp struct {
	Data struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"data"`
}

// POST chat/decision
type DecisionResp struct {
	Data struct {
		Content string `json:"content"`
	} `json:"data"`
}
