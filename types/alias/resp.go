package aliastypes

// GET /aliaes
type GetAliasResp struct {
	Aliaes []*GetAliasRespUnit `json:"aliaes"`
}

type GetAliasRespUnit struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Cmd         string `json:"cmd"`
}
