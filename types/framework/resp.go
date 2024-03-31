package frameworktypes

import "github.com/sagayosa/sayo_utils/module"

// GET /module/role
type GetModuleByRoleResp struct {
	Modules []*module.Module `json:"modules"`
}
