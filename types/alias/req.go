package aliastypes

// POST /aliaes
type SetAliasReq struct {
	Key         string                 `json:"key"`
	Description string                 `json:"description"`
	Cmd         string                 `json:"cmd"`
	Root        string                 `json:"root"`
	Argvs       map[string]interface{} `json:"argvs"`
}

// GET /aliaes
type GetAliasReq struct{}

// POST /mapping/:alias
type MappingReq struct{}

// GET /frameworkaddr
type GetFrameworkAddrReq struct{}

// DELETE /alias
type DeleteAliasReq struct {
	Key string `json:"key"`
}

// POST /aliaes
type SetAliaesReq struct {
	Aliaes []*SetAliasReq `json:"aliaes"`
}

// Delete /aliaes
type DeleteAliaesReq struct {
	Keys []string `json:"keys"`
}
