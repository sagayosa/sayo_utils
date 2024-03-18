package sdk

type NewWindowReq struct {
	Theme    string      `json:"theme"`
	Url      string      `json:"url"`
	Frame    bool        `json:"frame"`
	Dragable bool        `json:"dragable"`
	Option   interface{} `json:"option"`
}

type RegisterHotKeyReq struct {
	Identifier string `json:"identifier"`
	Key        string `json:"key"`
	Url        string `json:"url"`
}

type AICompletionsReq struct {
	Model       string     `json:"model"`
	Messages    []Messages `json:"messages"`
	Temperature int        `json:"temperature"`
	Stream      bool       `json:"stream"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AIDecisionResp struct {
	Root  string                 `json:"root"`
	Argvs map[string]interface{} `json:"argvs"`
}

type CoreVoiceCommandReq struct {
	Path string `json:"path"`
}

type GetModuleByRoleReq struct {
	Role string `json:"role"`
}
