package aitypes

// chat/completions
type CompletionsResp struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// POST chat/decision
type DecisionResp struct {
	Content string `json:"content"`
}
