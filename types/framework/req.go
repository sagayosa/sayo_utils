package frameworktypes

// GET /module/role"
type GetModuleByRoleReq struct {
	Role string `json:"role"`
}

// POST /proxy/plugin
type ProxyPluginReq struct {
	Root  string                 `json:"root"`
	Argvs map[string]interface{} `json:"argvs"`
}
