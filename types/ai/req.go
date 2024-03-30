package aitypes

// POST chat/completions
type CompletionsReq struct {
	Model       string     `json:"model"`
	Messages    []Messages `json:"messages"`
	Temperature int        `json:"temperature"`
	Stream      bool       `json:"stream"`
}
type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// PUT /config
type UpdateConfigReq struct {
	CompletionsURL string `json:"completions_url"`
	APIKey         string `json:"api_key"`
	Model          string `json:"model"`
}

// POST chat/decision
type DecisionReq struct {
	Content string `json:"content"`
}

// GET /config
type GetConfigReq struct{}
