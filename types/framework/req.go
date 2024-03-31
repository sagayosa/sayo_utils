package frameworktypes

// GET /module/role"
type GetModuleByRoleReq struct {
	Role string `json:"role"`
}

// POST /proxy/plugin
type ProxyPluginReq struct {
	Content string `json:"content"`
}
