package desktoptypes

// GET /fileselector
type OpenFileSelectorReq struct {
}

// POST /window
type NewWindowReq struct {
	Theme    string      `json:"theme"`
	Url      string      `json:"url"`
	Frame    bool        `json:"frame"`
	Dragable bool        `json:"dragable"`
	Option   interface{} `json:"option"`
}
