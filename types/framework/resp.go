package frameworktypes

import "github.com/sagayosa/sayo_utils/module"

// GET /module/role"
type GetModuleByRoleResp struct {
	Data []*module.Module `json:"data"`
}
