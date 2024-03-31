package opentypes

// POST /open
type OpenReq struct {
	Identifier string `json:"identifier"`
	Accurate   bool   `json:"accurate"`
}

type Identifier struct {
	Identifier  string `json:"identifier"`
	ExecBin     string `json:"bin"`
	Description string `json:"description"`
	Allow       bool   `json:"allow"`
}

// POST /identifiers
type RegisterIdentifierReq struct {
	Identifiers []*Identifier `json:"identifiers"`
}

// DELETE /identifiers
type UnRegisterIdentifierReq struct {
	Identifiers []string `json:"identifiers"`
}

// POST /identifier/allow
type AllowIdentifierReq struct {
	Identifier string `json:"identifier"`
}

// DELETE /identifier/allow
type DisAllowIdentifierReq struct {
	Identifier string `json:"identifier"`
}

// GET /identifiers
type GetIdentifiersReq struct{}

// GET /fileselector
type OpenFileSelectorReq struct{}

// PUT /identifiers
type UpdateIdentifiersReq struct {
	Identifiers []*Identifier `json:"identifiers"`
}
